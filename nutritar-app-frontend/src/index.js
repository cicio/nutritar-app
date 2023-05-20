import React from 'react';
import ReactDOM from 'react-dom/client';
import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import App from './App';
import ErrorPage from './components/ErrorPage';
import Foods from './components/Foods';
import Home from './components/Home';
import Nutrients from './components/Nutrients';
import EditFood from './components/EditFood';
import ManageMeals from './components/ManageMeals';
import GraphQL from './components/GraphQL';
import Login from './components/Login';
import Food from './components/Food'



const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,
    errorElement: <ErrorPage />,
    children: [
      {index: true, element: <Home />},
      {path: "/foods", element: <Foods />},
      {path: "/foods/:id", element: <Food />},
      {path: "/nutrients", element: <Nutrients />},
      {path: "/admin/food/0", element: <EditFood />},
      {path: "/manage-meals", element: <ManageMeals />},
      {path: "/graphql", element: <GraphQL />},
      {path: "/login", element: <Login />},
    ]

  }
])


const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
     <RouterProvider router={router} />
  </React.StrictMode>
);

