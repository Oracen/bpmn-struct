package common

import (
	"testing"

	testhelpers "github.com/Oracen/bpmn-struct/test_helpers"
)

var (
	testCorrExpression = CreateFormalExpression("id", "body", CreateItemDefinition("id", ITEM_KIND_Information))
	testCorrMessage    = CreateMessage("id")
	refExpressions     = []CorrelationPropertyRetrievalExpression{
		CreateCorrelationPropertyRetrievalExpression("id", testCorrExpression, testCorrMessage),
	}
)

func TestCorrelationKey(t *testing.T) {
	name, fn := testhelpers.CreateDefaultIsValid(
		t,
		CreateCorrelationKey("id"),
	)
	t.Run(name, fn)
}

func TestCorrelationPropertyRetrievalExpression(t *testing.T) {

	name, fn := testhelpers.CreateDefaultIsValid(
		t,
		CreateCorrelationPropertyRetrievalExpression("id", testCorrExpression, testCorrMessage),
	)
	t.Run(name, fn)
}

func TestCorrelationProperty(t *testing.T) {

	name, fn := testhelpers.CreateDefaultIsValid(
		t,
		CreateCorrelationProperty("id", refExpressions),
	)
	t.Run(name, fn)
}

func TestCorrelationSubscription(t *testing.T) {
	name, fn := testhelpers.CreateDefaultIsValid(
		t,
		CreateCorrelationSubscription("id", CreateCorrelationKey("id_too")),
	)
	t.Run(name, fn)
}

func TestCorrelationPropertyBinding(t *testing.T) {
	name, fn := testhelpers.CreateDefaultIsValid(
		t,
		CreateCorrelationPropertyBinding(
			"id",
			testCorrExpression,
			CreateCorrelationProperty("id2", refExpressions),
		),
	)
	t.Run(name, fn)
}
