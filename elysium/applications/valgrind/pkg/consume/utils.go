package consume

import "elysium.com/shared/types"

func divideBatchIntoBatchesBySubject(primo Batch) []Batch {

	batchMap := make(map[string][]types.Prototype)
	for _, prototype := range primo.prototypes {
		batchMap[prototype.GetSubject()] = append(batchMap[prototype.GetSubject()], prototype)
	}

	batches := make([]Batch, 0)

	for subject := range batchMap {
		batches = append(
			batches, Batch{
				subject:    subject,
				prototypes: batchMap[subject],
			},
		)
	}

	return batches
}
