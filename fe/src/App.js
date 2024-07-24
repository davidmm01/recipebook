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
      descriptorsSelected: {},
      recipes: [],
      DataisLoaded: false,
    };
    this.cuisineSelectHandler = this.cuisineSelectHandler.bind(this);
    this.descriptorsSelectHandler = this.descriptorsSelectHandler.bind(this);
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

  getRecipes(latestCusineSelection, latestDescriptorsSelection) {
    const searchParams = new URLSearchParams();
    if (latestCusineSelection != "") {
      searchParams.append("cuisine", latestCusineSelection);
    }

    var values = Object.keys(latestDescriptorsSelection).map(function (key) {
      return latestDescriptorsSelection[key].Name;
    });
    var desciptorNamesCommaDelimited = values.join(",");
    if (desciptorNamesCommaDelimited != "") {
      searchParams.append("descriptors", desciptorNamesCommaDelimited);
    }
    console.log("searchParams.toString():", searchParams.toString());
    fetch("http://localhost:8080/recipes?" + searchParams.toString())
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
    this.getRecipes(cuisineSelection, this.state.descriptorsSelected);
  }

  descriptorsSelectHandler(descriptorsSelection) {
    this.setState({ descriptorsSelected: descriptorsSelection });
    console.log("descriptors selected in the App:", descriptorsSelection);
    this.getRecipes(this.state.cuisineSelected, descriptorsSelection);
  }

  render() {
    const { DataisLoaded, recipes, cuisineSelected, descriptorsSelected } =
      this.state;
    if (!DataisLoaded)
      return (
        <div>
          <h1> Loading </h1>
        </div>
      );

    return (
      <div>
        <div className="grid-container">
          <Cuisines selectedCuisine={this.cuisineSelectHandler} />
          <Descriptors selectedDescriptors={this.descriptorsSelectHandler} />
        </div>
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
