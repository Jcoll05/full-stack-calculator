import './App.css'
import Calculator from './components/Calculator'

function App() {
  return (
    <main className="app">
      <header className="app-header">
        <div>
          <span className="app-eyebrow">FULL-STACK</span>
          <h1>Calculator</h1>
        </div>

        <span className="app-status">ONLINE</span>
      </header>

      <Calculator />

    </main>
  )
}

export default App