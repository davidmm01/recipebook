import React from "react";

// keep this as separate to cuisines since both will need to be different shortly
class Descriptors extends React.Component {
  // Constructor
  constructor(props) {
    super(props);

    this.state = {
      descriptors: [],
      DataisLoaded: false,
      selectedDescriptors: {},
    };
  }

  toggleSelectedDescriptorState = (key) => {
    // on re-click of already selected descriptor, clear selection
    if (this.state.selectedDescriptors[key]) {
      delete this.state.selectedDescriptors[key];
      this.state.descriptors[key]["IsSelected"] = false;
      this.setState(
        {
          selectedDescriptors: this.state.selectedDescriptors,
          descriptors: this.state.descriptors,
        },
        () => console.log(this.state)
      );
    } else {
      this.state.selectedDescriptors[key] = this.state.descriptors[key];
      this.state.descriptors[key]["IsSelected"] = true;
      this.setState(
        {
          selectedDescriptors: this.state.selectedDescriptors,
          descriptors: this.state.descriptors,
        },
        console.log(this.state)
      );
    }
    this.props.selectedDescriptors(this.state.selectedDescriptors);
  };

  // ComponentDidMount is used to
  // execute the code
  componentDidMount() {
    fetch("http://localhost:8080/descriptors")
      .then((res) => res.json())
      .then((json) => {
        this.setState({
          descriptors: json,
          DataisLoaded: true,
        });
      });
  }

  render() {
    const { DataisLoaded, descriptors } = this.state;
    if (!DataisLoaded)
      return (
        <div>
          <h1> Pleses wait some time.... </h1>
        </div>
      );

    return (
      <div className="container">
        <h2>Descriptors</h2>
        <div className="container grid-container">
          {descriptors.map((descriptor, i) => (
            <div>
              <ol>
                <button
                  key={i}
                  name={descriptor.Name}
                  onClick={() => this.toggleSelectedDescriptorState(i)}
                  className={
                    descriptor.IsSelected
                      ? "grid-item customButton active"
                      : "grid-item customButton"
                  }
                >
                  {descriptor.Name} ({descriptor.Count})
                </button>
              </ol>
            </div>
          ))}
        </div>
      </div>
    );
  }
}

export default Descriptors;
