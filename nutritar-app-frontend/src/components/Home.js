import React from "react";
import { Link } from "react-router-dom";

import Mainpicture from './../images/homepic.jpg';

const Home = () => {

    return(
        <>
        <div className="text-center">
            <h2>Find foods to consume today</h2>
            <hr />
            <Link to="/foods">
                <img src={Mainpicture} alt="food nutrients"></img>
            </Link>
        </div>
        </>
    );
};

export default Home;