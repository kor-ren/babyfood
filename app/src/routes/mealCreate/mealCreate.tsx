import { useMutation } from "@apollo/client"
import { SubmitHandler, useForm } from "react-hook-form"
import { CREATE_MEAL, GET_MEALS } from "../../queries"
import { NewMeal } from "../../__generated__/graphql"
import { useNavigate } from "react-router-dom"

import './createMeal.css'


export default function MealCreate() {
    const navigation = useNavigate()

    const {
        register,
        handleSubmit,
    } = useForm<NewMeal>()

    const [createMeal, { loading }] = useMutation(CREATE_MEAL, {
        refetchQueries: [
            GET_MEALS
        ]
    })

    const onSubmit: SubmitHandler<NewMeal> = (data) => {
        createMeal({
            variables: {
                input: {
                    ...data,
                    rating: !!data.rating ? data.rating : undefined
                }
            },
            onCompleted: () => {
                navigation("/")
            }
        })
    }

    return (
        <form onSubmit={handleSubmit(onSubmit)}>
            <input placeholder="Name" {...register("name")} />
            <input placeholder="Image" {...register("image")} />
            <input placeholder="Rating" type="number" {...register("rating")} />

            <input type="submit" disabled={loading} />
        </form>
    )
}