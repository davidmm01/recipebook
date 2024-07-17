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
    // There is surely a much better way of building up query params with a popular lib,
    // look into this.
    // Or maybe query params are just kinda gross and could you a body with request
    var queryParam = "";
    if (
      latestCusineSelection != "" ||
      Object.keys(latestDescriptorsSelection).length != 0
    ) {
      queryParam += "?";
      if (latestCusineSelection != "") {
        queryParam += "cuisine=" + latestCusineSelection;
      }

      console.log(1);
      if (Object.keys(latestDescriptorsSelection).length != 0) {
        console.log(2);
        if (queryParam != "?") {
          // i.e. it has other query in it
          queryParam += "&";
        }
        queryParam += "descriptors=";
        for (const [key, value] of Object.entries(latestDescriptorsSelection)) {
          // Need to fix this, we now have spaces in the query param. Should probably be passing ids around instead
          queryParam += value.Name + ",";
          // console.log(key, value);
        }
        queryParam = queryParam.slice(0, -1);
      }
    }
    console.log("FINAL QUERY PARAM:", queryParam);
    // fetch("http://localhost:8080/recipes" + queryParam)
    //   .then((res) => res.json())
    //   .then((json) => {
    //     this.setState({
    //       recipes: json,
    //       DataisLoaded: true,
    //     });
    //   });
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
          <h1> Pleses wait some time.... </h1>
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
