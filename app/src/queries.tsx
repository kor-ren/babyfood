import {gql} from './__generated__/gql'


export const GET_MEALS = gql(/* GraphQL */ `
    query getMeals {
        meals {
            id
            name
            rating
        }
    }
    `
)

export const CREATE_MEAL = gql(/* GraphQL */ `
    mutation createMeal($input: NewMeal!) {
        createMeal(input: $input) {
            id
            name
        }
    }
    `)