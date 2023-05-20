import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import foodAndi from '../foodandi.json';

const Food = () => {
    const [food, setFood] = useState([]);
    let {id }= useParams();
    
    useEffect(() => {
        let myFood = {
            id: 1,
            name: "Kale",
            gombsCategory: "greens",
            foodGroup: "vegetable"
        }
        setFood(myFood);
    }, [id])

    return(

        <div>
            <h2>Food: {food.name}</h2>
            <small><em>ANDI Score: {foodAndi[food.name]}</em></small>
            <hr />
            <p>G-BOMBS Category: {food.gombsCategory}</p>
            <p>Food Group: {food.foodGroup}</p>
        </div>
    );
};

export default Food;