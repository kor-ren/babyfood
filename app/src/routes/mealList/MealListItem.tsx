import { Meal } from "../../__generated__/graphql"

import './mealListItem.css'

export type MealListITemProps = {
    meal: Pick<Meal, 'id' | 'name' | 'image'>
}

export default function MealListItem({ meal }: MealListITemProps) {

    return (
        <figure key={meal.id}>
            <img src="https://thumbs.dreamstime.com/b/abstract-background-red-metallic-poster-vector-illustration-145174208.jpg" />
            <figcaption>
                {meal.name}
            </figcaption>
        </figure>
    )
}