import React, { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import foodAndi from '../foodandi.json'
//import kale from '../images/freshkale.jpg'

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
                        <td>Description</td>
                        <td>Image</td>
                        <td>ANDI Score</td>
                    </tr>
                </thead>
                <tbody>
                    {foods.map((f) => (
                        <tr key={f.id}>
                            <td>
                                <Link to={`/foods/${f.id}`}>
                                    {f.food_name}
                                </Link>
                            </td>
                            <td>{f.food_description}</td>
                            <td>{f.food_image}</td>
                            <td>{foodAndi[f.name]}</td>
                        </tr>
                    ))}
                </tbody>
            </table>
            
        </div>
    );
};

export default Foods;
