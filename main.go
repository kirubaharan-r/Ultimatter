package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/ghodss/yaml"
)

type time struct {
	Start string
	End   string
	Type1 string
	Type2 string
}

func main() {
	sat := time{
		End:   "2023-04-01",
		Start: "2023-03-01",
		Type1: "DIMENSION",
		Type2: "TAG",
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-west-2"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	ce := costexplorer.NewFromConfig(cfg)
	groups := []types.GroupDefinition{
		{
			Type: types.GroupDefinitionType(sat.Type1),
			Key:  aws.String("SERVICE"),
		},
		{
			Type: types.GroupDefinitionType(sat.Type2),
			Key:  aws.String("Environment"),
		},
	}
	input := &costexplorer.GetCostAndUsageInput{
		Granularity: "MONTHLY",
		TimePeriod:  &types.DateInterval{End: &sat.End, Start: &sat.Start},
		Metrics:     []string{"UnblendedCost", "BlendedCost"},
		GroupBy:     groups,
	}

	resp, err := ce.GetCostAndUsage(context.TODO(), input)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	//fmt.Println(resp.ResultMetadata)
	jsonOutput, _ := json.Marshal(resp)
	yamlout, _ := yaml.Marshal(resp)
	//yamlout, _ := yaml.JSONToYAML(jsonOutput)

	fmt.Println(string(jsonOutput))
	fmt.Println(string(yamlout))

}
