import { useCallback } from 'react'
import './App.css'
import { useMutation, useQuery } from '@apollo/client'
import { CREATE_MEAL, GET_MEALS } from './queries'

function App() {

  const { data, loading, error, refetch } = useQuery(GET_MEALS)

  const [createMeal] = useMutation(CREATE_MEAL)

  const createMealHandler = useCallback(() => {
    createMeal({
      variables: {
        input: {
          name: "Test",
          rating: 2
        }
      }
    }).finally(() => refetch())
  }, [createMeal, refetch])

  if (loading) {
    return <div>Loading...</div>
  }

  if (error) {
    return <pre>{JSON.stringify(error, null, 2)}</pre>
  }

  if (!data) {
    return <div>No data</div>
  }


  return (
    <>
      <table>
        <tr>
          <th>Name</th>
          <th>Rating</th>
        </tr>
      {data.meals.map(m => (
        <tr key={m.id}>
          <td>{m.name}</td>
          <td>{m.rating}</td>
        </tr>
      ))}
    </table>
    <button onClick={createMealHandler}>CreateMeal</button>
    </>
  )
}

export default App
