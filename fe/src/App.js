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
      recipes: [],
      DataisLoaded: false,
    };
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

  render() {
    const { DataisLoaded, recipes } = this.state;
    if (!DataisLoaded)
      return (
        <div>
          <h1> Pleses wait some time.... </h1>
        </div>
      );

    return (
      <div>
        <Cuisines />
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
