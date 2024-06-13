import {gql} from './__generated__/gql'


export const GET_MEALS = gql(/* GraphQL */ `
    query getMeals {
        meals {
            id
            name
            rating
            image
        }
    }
    `
)

export const GET_MEAL_BY_ID = gql(/* GraphQL */ `
    query getMealById($id: ID!) {
        meal(id: $id) {
            id
            name
            rating
            image
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

