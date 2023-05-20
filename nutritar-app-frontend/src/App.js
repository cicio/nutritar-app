import React, { useState } from "react";

import { Link, Outlet, useNavigate } from "react-router-dom";
import Alert from "./components/Alert";

const App = () => {
  const [jwtToken, setJwtToken] = useState("");
  const [alertMessage, setAlertMessage] = useState("");
  const [alertClassName, setAlertClassName]= useState("d-none");

  const navigate = useNavigate();

  const logOut = () => {
    setJwtToken("");
    navigate("/login")
  }



  return (
    <div className="container">
      <div className="row">
        <div className="col">
          <h1 className="mt-3">Be Nutritarian!</h1>
        </div>
        <div className="col text-end">
          {jwtToken === ""
          ? <Link to="/login"><span className="badge bg-success mt-3">Login</span></Link>
          : <a href="#!" onClick={logOut}><span className="badge bg-danger">Logout</span></a>
          }
        </div>
        <hr className="mb-3" />
      </div>
      
      <div className="row">
        <div className="col-md-2">
          <nav>
            <div className="list-group">
              <Link to="/" className="list-group-item list-group-item-action">Home</Link>
              <Link to="/foods" className="list-group-item list-group-item-action">Foods</Link>
              <Link to="/nutrients" className="list-group-item list-group-item-action">Nutrients</Link>
              {jwtToken !== "" &&
                <>
                <Link to="admin/food/0" className="list-group-item list-group-item-action">Add Food</Link>
                <Link to="/manage-meals" className="list-group-item list-group-item-action">Manage Meals</Link>
                <Link to="/graphql" className="list-group-item list-group-item-action">GraphQL</Link>
                </>
              }

            </div>
          </nav>
        </div>
        <div className="col-md-10">
          <Alert 
            message={alertMessage}
            className={alertClassName}
          />
          <Outlet context={{
            jwtToken, 
            setJwtToken,
            setAlertMessage,
            setAlertClassName
          }}/>
        </div>
      </div>
    </div>
  );
}

export default App;
