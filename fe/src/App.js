// App.js
import React from "react";
import "./App.css";
import RecipeCard from "./components/RecipeCard";
import Cuisines from "./components/Cuisines";
import Descriptors from "./components/Descriptors";

class App extends React.Component {
  // Constructor
  constructor(props) {
    super(props);

    this.state = {
      cuisineSelected: "",
      recipes: [],
      DataisLoaded: false,
    };
    this.cuisineSelectHandler = this.cuisineSelectHandler.bind(this);
  }

  // ComponentDidMount is used to
  // execute the code
  componentDidMount() {
    fetch("http://localhost:8080/recipes")
      .then((res) => res.json())
      .then((json) => {
        this.setState({
          recipes: json,
          DataisLoaded: true,
        });
      });
  }

  cuisineSelectHandler(cuisineSelection) {
    this.setState({ cuisineSelected: cuisineSelection });
    console.log("cusuine selected in the App:", cuisineSelection);
  }

  render() {
    const { DataisLoaded, recipes, cuisineSelected } = this.state;
    if (!DataisLoaded)
      return (
        <div>
          <h1> Pleses wait some time.... </h1>
        </div>
      );

    return (
      <div>
        <Cuisines selectedCuisine={this.cuisineSelectHandler} />
        <Descriptors />
        <h2>Recipes</h2>
        <div className="container">
          {recipes.map((recipe) => (
            <div className="recipe">
              <ol>
                <RecipeCard name={recipe.Name} />
              </ol>
            </div>
          ))}
        </div>
      </div>
    );
  }
}

export default App;
