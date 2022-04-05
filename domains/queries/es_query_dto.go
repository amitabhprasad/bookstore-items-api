package queries

type EsQuery struct {
	Equals []FieldValue `json:"equals"`
}

type FieldValue struct {
	Field string      `json:"field"`
	Value interface{} `json:"value"`
}

/**
	Sample:
	{
		"equals":
		[
			{
				"field":"status"
				"value":"pending"
			},
			"equals":{
				"field":"seller"
				"value":"7"
			}
		],
		"any_equals":
		[
			{
				"field":"status"
				"value":"pending"
			},
			"equals":{
				"field":"seller"
				"value":"7"
			}
		],
	}
**/
