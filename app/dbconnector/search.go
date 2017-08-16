package dbconnector

import (
	"encoding/json"
	"io"
	"context"
	"os"
	"log"
	"revelSample/app/models"
	elastic "gopkg.in/olivere/elastic.v5"
)

func Search(name string) ([]*models.ViewSearchResult, error) {

	// elasticSearch接続
	client, err := elastic.NewClient(elastic.SetURL(os.Getenv("ELASTICSEARCH_URL")), elastic.SetSniff(false))
	if err != nil {
		log.Printf("elasticSearch接続失敗")
		return nil, err
	}

	// クエリ発行
	var svc *elastic.ScrollService
	if len(name) != 0 {
		query := elastic.NewTermQuery("Name", name)
		svc = client.Scroll().
			Index("employee_info").
			Type("info").
			Query(query).
			Sort("number", true).
			Size(5000)
	} else {
		svc = client.Scroll().
			Index("employee_info").
			Type("info").
			Sort("Number", true).
			Size(5000)
	}

	// 検索結果取得
	var results []*models.ViewSearchResult
	for {
		res, err := svc.Do(context.TODO())
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("elasticSearch結果取得失敗 : %s", err.Error())
			return nil, err
		}
		if res == nil {
			log.Printf("expected results != nil; got nil")
			return nil, err
		}

		for _, hit := range res.Hits.Hits {
			var s models.DBRowData
			err := json.Unmarshal(*hit.Source, &s)
			if err != nil {
				// deserialization failed
				log.Printf("Unmarshal失敗")
				return nil, err
			}

			result := new(models.ViewSearchResult)
			result.SetView(s.Number, s.Name, s.Affiliation, s.Responsible, s.Mail,
				s.Tel, s.Entry_month, s.Birthday)
			results = append(results, result)
		}
	}
	
	return results, nil
}