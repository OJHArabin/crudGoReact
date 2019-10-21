

class App extends React.Component {
    render() {
      
        return (<Home />);
      
    }
  }
  class Home extends React.Component {
   
    render() {
      return (
        
        <div className="container">
          <div className="col-xs-8 col-xs-offset-2 jumbotron text-center">
            <h1>CRUD APP</h1>
            <p>Sign in to get access Or Register</p>
            <a onClick={this.authenticate} className="btn btn-primary btn-lg btn-login btn-block">Sign In</a>
            <a onClick={this.register} className="btn btn-primary btn-lg btn-login btn-block">Register</a>
          </div>
        </div>
      );
    }
  }

  class SignIn extends React.Component {
    
    render() {
      
      return (
        <h1>Hello, world!</h1>
      );
    } 
  }

  
   

  ReactDOM.render(<App />, document.getElementById('app'));