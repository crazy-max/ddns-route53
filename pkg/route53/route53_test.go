package route53

import (
	"context"
	"io"
	"net"
	"net/http"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	awsr53 "github.com/aws/aws-sdk-go-v2/service/route53"
	awsr53types "github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type httpClientFunc func(*http.Request) (*http.Response, error)

func (f httpClientFunc) Do(req *http.Request) (*http.Response, error) {
	return f(req)
}

func newTestClient(t *testing.T, httpClient httpClientFunc) *Client {
	t.Helper()

	return &Client{
		ctx: context.Background(),
		client: awsr53.New(awsr53.Options{
			Region:      "us-east-1",
			Credentials: credentials.NewStaticCredentialsProvider("test-access-key", "test-secret-key", ""),
			HTTPClient:  httpClient,
			Retryer:     aws.NopRetryer{},
		}, func(o *awsr53.Options) {
			o.BaseEndpoint = aws.String("https://route53.test")
		}),
		hostedZoneID: "ZTEST123",
	}
}

func xmlResponse(req *http.Request, status int, body string) *http.Response {
	return &http.Response{
		StatusCode: status,
		Header: http.Header{
			"Content-Type": []string{"application/xml"},
		},
		Body:       io.NopCloser(strings.NewReader(body)),
		Request:    req,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
	}
}

func TestListRecordsPaginates(t *testing.T) {
	t.Parallel()

	requests := 0
	client := newTestClient(t, func(req *http.Request) (*http.Response, error) {
		requests++
		require.Equal(t, http.MethodGet, req.Method)
		require.Equal(t, "/2013-04-01/hostedzone/ZTEST123/rrset", req.URL.Path)

		switch requests {
		case 1:
			assert.Empty(t, req.URL.Query().Get("name"))
			assert.Empty(t, req.URL.Query().Get("type"))
			assert.Empty(t, req.URL.Query().Get("identifier"))
			return xmlResponse(req, http.StatusOK, `
<ListResourceRecordSetsResponse xmlns="https://route53.amazonaws.com/doc/2013-04-01/">
  <IsTruncated>true</IsTruncated>
  <MaxItems>100</MaxItems>
  <NextRecordName>ipv6.example.com.</NextRecordName>
  <NextRecordType>AAAA</NextRecordType>
  <NextRecordIdentifier>next-page</NextRecordIdentifier>
  <ResourceRecordSets>
    <ResourceRecordSet>
      <Name>ipv4.example.com.</Name>
      <Type>A</Type>
      <TTL>300</TTL>
      <ResourceRecords>
        <ResourceRecord>
          <Value>203.0.113.10</Value>
        </ResourceRecord>
      </ResourceRecords>
    </ResourceRecordSet>
  </ResourceRecordSets>
</ListResourceRecordSetsResponse>`), nil
		case 2:
			assert.Equal(t, "ipv6.example.com.", req.URL.Query().Get("name"))
			assert.Equal(t, "AAAA", req.URL.Query().Get("type"))
			assert.Equal(t, "next-page", req.URL.Query().Get("identifier"))
			return xmlResponse(req, http.StatusOK, `
<ListResourceRecordSetsResponse xmlns="https://route53.amazonaws.com/doc/2013-04-01/">
  <IsTruncated>false</IsTruncated>
  <MaxItems>100</MaxItems>
  <ResourceRecordSets>
    <ResourceRecordSet>
      <Name>ipv6.example.com.</Name>
      <Type>AAAA</Type>
      <TTL>300</TTL>
      <ResourceRecords>
        <ResourceRecord>
          <Value>2001:db8::10</Value>
        </ResourceRecord>
      </ResourceRecords>
    </ResourceRecordSet>
  </ResourceRecordSets>
</ListResourceRecordSetsResponse>`), nil
		default:
			t.Fatalf("unexpected extra request %d", requests)
			return nil, nil
		}
	})

	records, err := client.ListRecords()
	require.NoError(t, err)
	require.Len(t, records, 2)
	assert.Equal(t, "ipv4.example.com.", aws.ToString(records[0].Name))
	assert.Equal(t, awsr53types.RRTypeA, records[0].Type)
	assert.Equal(t, "203.0.113.10", aws.ToString(records[0].ResourceRecords[0].Value))
	assert.Equal(t, "ipv6.example.com.", aws.ToString(records[1].Name))
	assert.Equal(t, awsr53types.RRTypeAaaa, records[1].Type)
	assert.Equal(t, "2001:db8::10", aws.ToString(records[1].ResourceRecords[0].Value))
	assert.Equal(t, 2, requests)
}

func TestListRecordsWrapsFetchErrors(t *testing.T) {
	t.Parallel()

	client := newTestClient(t, func(req *http.Request) (*http.Response, error) {
		require.Equal(t, http.MethodGet, req.Method)
		return xmlResponse(req, http.StatusBadRequest, `
<ErrorResponse>
  <Error>
    <Code>InvalidInput</Code>
    <Message>bad request</Message>
  </Error>
  <RequestId>req-123</RequestId>
</ErrorResponse>`), nil
	})

	records, err := client.ListRecords()
	require.Nil(t, records)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to fetch records")
	assert.Contains(t, err.Error(), "bad request")
}

func TestUpdateSendsHostedZoneAndChangeBatch(t *testing.T) {
	t.Parallel()

	var body string
	client := newTestClient(t, func(req *http.Request) (*http.Response, error) {
		require.Equal(t, http.MethodPost, req.Method)
		require.Equal(t, "/2013-04-01/hostedzone/ZTEST123/rrset", req.URL.Path)

		payload, err := io.ReadAll(req.Body)
		require.NoError(t, err)
		body = string(payload)

		return xmlResponse(req, http.StatusOK, `
<ChangeResourceRecordSetsResponse xmlns="https://route53.amazonaws.com/doc/2013-04-01/">
  <ChangeInfo>
    <Id>/change/C123456</Id>
    <Comment>updated from test</Comment>
    <Status>PENDING</Status>
    <SubmittedAt>2026-04-19T20:00:00Z</SubmittedAt>
  </ChangeInfo>
</ChangeResourceRecordSetsResponse>`), nil
	})

	resp, err := client.Update([]awsr53types.Change{
		{
			Action: awsr53types.ChangeActionUpsert,
			ResourceRecordSet: &awsr53types.ResourceRecordSet{
				Name: aws.String("ipv4.example.com."),
				Type: awsr53types.RRTypeA,
				TTL:  aws.Int64(300),
				ResourceRecords: []awsr53types.ResourceRecord{
					{Value: aws.String("203.0.113.42")},
				},
			},
		},
	}, "updated from test")
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, resp.ChangeInfo)

	assert.Contains(t, body, "<Comment>updated from test</Comment>")
	assert.Contains(t, body, "<Name>ipv4.example.com.</Name>")
	assert.Contains(t, body, "<Type>A</Type>")
	assert.Contains(t, body, "<TTL>300</TTL>")
	assert.Contains(t, body, "<Value>203.0.113.42</Value>")
	assert.Equal(t, "/change/C123456", aws.ToString(resp.ChangeInfo.Id))
	assert.Equal(t, awsr53types.ChangeStatusPending, resp.ChangeInfo.Status)
}

func TestRecordIPReturnsMatchingAddress(t *testing.T) {
	t.Parallel()

	client := &Client{}
	records := []awsr53types.ResourceRecordSet{
		{
			Name: aws.String("ignored.example.com."),
			Type: awsr53types.RRTypeA,
			ResourceRecords: []awsr53types.ResourceRecord{
				{Value: aws.String("198.51.100.10")},
			},
		},
		{
			Name: aws.String("match.example.com."),
			Type: awsr53types.RRTypeAaaa,
			ResourceRecords: []awsr53types.ResourceRecord{
				{Value: aws.String("2001:db8::42")},
			},
		},
	}

	ip, err := client.RecordIP(records, aws.String("match.example.com."), awsr53types.RRTypeAaaa)
	require.NoError(t, err)
	assert.True(t, net.ParseIP("2001:db8::42").Equal(ip))

	ip, err = client.RecordIP(records, aws.String("missing.example.com."), awsr53types.RRTypeA)
	require.NoError(t, err)
	assert.Nil(t, ip)
}
