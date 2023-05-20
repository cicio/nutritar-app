import React, { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import foodAndi from '../foodandi.json'

const Foods = () => {
    const [foods, setFoods] = useState([]);

    useEffect(() => {
        let foodsList = [
            {
            id: 1,
            name: "Kale",
            gombsCategory: "greens",
            foodGroup: "vegetable",
            },
            {
            id: 2,
            name: "Watercress",
            gombsCategory: "greens",
            foodGroup: "vegetable",
            },
            {
            id: 3,
            name: "Apple",
            gombsCategory: "fruit",
            foodGroup: "fruit",
            }

        ];
        setFoods(foodsList)
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
                        <tr key={f.id}>
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
