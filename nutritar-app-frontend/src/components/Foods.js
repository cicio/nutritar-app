import React, { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import foodAndi from '../foodandi.json'

const Foods = () => {
    const [foods, setFoods] = useState([]);

    useEffect(() => {
       const headers = new Headers();
       headers.append("Content-Type", "application/json");

       const requestOptions = {
        method: "GET",
        headers: headers,
       }

       fetch(`http://localhost:8080/foods`, requestOptions)
            .then((response) => response.json())
            .then((data) => {
                setFoods(data);
            })
            .catch(err => {
                console.log(err);
            })


    },[]);

    return(
        <div>
            <h2>Foods</h2>
            <hr />
            <table className="table table-hover table-striped">
                <thead>
                    <tr>
                        <td>Food</td>
                        <td>ANDI Score</td>
                        <td>GOMBS Category</td>
                        <td>Food Group</td>
                    </tr>
                </thead>
                <tbody>
                    {foods.map((f) => (
                        <tr key={f.FoodID}>
                            <td>
                                <Link to={`/foods/${f.id}`}>
                                    {f.name}
                                </Link>
                            </td>
                            <td>{foodAndi[f.name]}</td>
                            <td>{f.gombsCategory}</td>
                            <td>{f.foodGroup}</td>
                        </tr>
                    ))}
                </tbody>
            </table>
            
        </div>
    );
};

export default Foods;
