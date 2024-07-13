import React from "react";
import "./cuisines.css";

// used the following to help make the button group
// https://dev.to/ramonak/react-how-to-create-a-custom-button-group-component-in-5-minutes-3lfd

class Cuisines extends React.Component {
  // Constructor
  constructor(props) {
    super(props);

    this.state = {
      cuisines: [],
      DataisLoaded: false,
      selectedCuisineKey: -1,
      selectedCuisine: "",
    };
  }

  setSelectedCuisineState = (key) => {
    // on re-click of already selected cuisine, clear selection
    if (this.state.selectedCuisineKey === key) {
      this.setState(
        {
          selectedCuisineKey: -1,
          selectedCuisine: "",
        },
        () => console.log(this.state)
      );
      this.props.selectedCuisine("");
    } else {
      this.setState(
        {
          selectedCuisineKey: key,
          selectedCuisine: this.state.cuisines[key].Name,
        },
        console.log(this.state)
      );
      this.props.selectedCuisine(this.state.cuisines[key].Name);
    }
  };

  // ComponentDidMount is used to
  // execute the code
  componentDidMount() {
    this.props.selectedCuisine(this.state.selectedCuisine);
    fetch("http://localhost:8080/cuisines")
      .then((res) => res.json())
      .then((json) => {
        this.setState({
          cuisines: json,
          DataisLoaded: true,
        });
      });
  }

  render() {
    const { DataisLoaded, cuisines, selectedCuisineKey } = this.state;
    if (!DataisLoaded)
      return (
        <div>
          <h1> Pleses wait some time.... </h1>
        </div>
      );

    return (
      <div>
        <h2>Cuisines</h2>
        <div className="container grid-container">
          {cuisines.map((cuisine, i) => (
            <div>
              <ol>
                <button
                  key={i}
                  name={cuisine.Name}
                  onClick={() => this.setSelectedCuisineState(i)}
                  className={
                    i === this.state.selectedCuisineKey
                      ? "grid-item customButton active"
                      : "grid-item customButton"
                  }
                >
                  {cuisine.Name} ({cuisine.Count})
                </button>
              </ol>
            </div>
          ))}
        </div>
      </div>
    );
  }
}

export default Cuisines;
