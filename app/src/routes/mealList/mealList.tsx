import { useQuery } from "@apollo/client"
import { GET_MEALS } from "../../queries"
import MealListItem from "./MealListItem"
import { Link } from "react-router-dom"

export default function MealList() {

    const { loading, error, data } = useQuery(GET_MEALS)


    if (loading) {
        return <div>Loading...</div>
    }

    if (error) {
        return <div>Could not load</div>
    }

    return (
        <>
            <div><Link to={'/create-meal'}>Neu</Link></div>
            <div className="list">
                {data?.meals.map(m => (
                    <MealListItem meal={m} />
                ))}
            </div>
        </>
    )
}