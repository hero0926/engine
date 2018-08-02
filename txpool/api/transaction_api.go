/*
 * Copyright 2018 It-chain
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package api

import (
	"log"

	"github.com/it-chain/engine/txpool"
)

type TransactionRepository interface {
	FindAll() ([]txpool.Transaction, error)
	Save(transaction txpool.Transaction) error
	Remove(id txpool.TransactionId)
	FindById(id txpool.TransactionId) (txpool.Transaction, error)
}

type TransactionApi struct {
	publisherId           string
	transactionRepository TransactionRepository
}

func NewTransactionApi(publisherId string, transactionRepository TransactionRepository) TransactionApi {
	return TransactionApi{
		publisherId:           publisherId,
		transactionRepository: transactionRepository,
	}
}

func (t TransactionApi) CreateTransaction(txData txpool.TxData) (txpool.Transaction, error) {

	transaction, err := txpool.CreateTransaction(t.publisherId, txData)

	if err != nil {
		log.Printf("fail to transaction: [%v]", err)
		return txpool.Transaction{}, err
	}

	err = t.transactionRepository.Save(transaction)

	return transaction, err
}

func (t TransactionApi) DeleteTransaction(id txpool.TransactionId) {

	t.transactionRepository.Remove(id)
}
