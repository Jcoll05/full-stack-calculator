import './App.css'
import Calculator from './components/Calculator'
import calculatorLogo from './assets/calculator-svgrepo-com.svg'

function App() {
  return (
    <main className="app">
      <header className="app-header">
        <div className="app-brand">
          <img
            className="app-logo"
            src={calculatorLogo}
            alt="Calculator logo"
          />
          <div>
            <span className="app-eyebrow">FULL-STACK</span>
            <h1>Calculator</h1>
          </div>
        </div>
        <span className="app-status">ONLINE</span>
      </header>

      <Calculator />

    </main>
  )
}

export default App