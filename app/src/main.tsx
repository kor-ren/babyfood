import React from 'react'
import ReactDOM from 'react-dom/client'
import './index.css'
import { ApolloClient, ApolloProvider, InMemoryCache } from '@apollo/client'
import { RouterProvider, createBrowserRouter } from 'react-router-dom'
import Root from './routes/root/root'
import MealList from './routes/mealList/mealList'
import MealCreate from './routes/mealCreate/mealCreate'


const client = new ApolloClient({
  uri: '/query',
  cache: new InMemoryCache()
})

const router = createBrowserRouter([
  {
    path: "/",
    element: <Root />,
    children: [
      {
        path: "/",
        element: <MealList />,
      },
      {
        path: "create-meal",
        element: <MealCreate />
      }
    ]
  }
])

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <ApolloProvider client={client}>
      <RouterProvider router={router} />
    </ApolloProvider>
  </React.StrictMode>,
)
