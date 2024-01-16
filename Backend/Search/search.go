package Search

import (
	"Backend/ent"
	"context"
	"log"
	"math"
	"sort"
)

func Ranking(data Website, client *ent.Client) []Data {
	// データからベクトルを生成する。
	wantVector := createVector(data)

	distanceList := make([]Data, 0, 10)
	// データベースから値を取得し、ベクトルにする。
	// 繰り返し処理でいい感じに。
	// 一旦仮にベクトルを作っときます。DBがんば
	dbData, err := client.Website.Query().All(context.Background())
	if err != nil {
		log.Fatalf(err.Error())
	}
	for _, db := range dbData {
		vector := createVector(Website{*db})
		// ベクトルの距離を計算する。
		var result float64
		for j, v := range wantVector {
			minusResult := v - vector[j]
			powResult := minusResult * minusResult

			result += powResult
		}

		// 距離をリストに代入
		result = math.Sqrt(result)
		distanceList = append(distanceList, Data{db.ID, result})
	}

	// 最小を求める
	sort.Slice(distanceList, func(i, j int) bool { return distanceList[i].Distance < distanceList[j].Distance })

	// jsonとして返す
	return distanceList
}

func createVector(data Website) []float64 {
	vector := make([]float64, 0, 7)
	if data.Adult != nil {
		vector = append(vector, *data.Adult)
	}
	if data.Flashy != nil {
		vector = append(vector, *data.Flashy)
	}
	if data.Beautiful != nil {
		vector = append(vector, *data.Beautiful)
	}
	if data.Like != nil {
		vector = append(vector, *data.Like)
	}
	if data.Smart != nil {
		vector = append(vector, *data.Smart)
	}

	return vector
}

type Website struct {
	ent.Website
}

type Data struct {
	Id       int
	Distance float64
}
