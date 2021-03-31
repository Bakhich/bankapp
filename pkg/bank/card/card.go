package card

import (
  "bank/pkg/bank/types"
  )
  
  func Total(card []types.Card) types.Money{
    
    sum := types.Money(0)

    for _, index := range card{
      inActive := index.Active
      sumBonus := index.Balance
      if inActive && sumBonus > 0 {
        sum += sumBonus
      }
        
    }
    return sum
  }