import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import foodAndi from '../foodandi.json';
//import kaleimg from './../images/freshkale.jpg'

const Food = () => {
    const [food, setFood] = useState([]);
    let {id }= useParams();
    
    useEffect(() => {
        let myFood = {
            id: 1,
            data_bank_id: "72119190",
            food_name: "Kale",
            food_description: "Kale fresh, raw",
            food_image: "./../images/freshkale.jpg",
            created_at: "2023-05-20",
            updated_at: "2023-05-20",
        }
        setFood(myFood);
    }, [id])

    return(

        <div>
            <h2>Food: {food.food_name}</h2>
            <small><em>ANDI Score: {foodAndi[food.food_name]}</em></small>
            <hr />
            <p>Description: {food.food_description}</p>
            <p>Image: {food.food_image}</p>
        </div>
    );
};

export default Food;