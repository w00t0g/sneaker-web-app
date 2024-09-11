import React from 'react'
import { Routes, Route, Link } from 'react-router-dom'
import Sneakers from './Sneakers'
import AddSneaker from './AddSneaker'
import SellSneaker from './SellSneaker'
import SoldSneakers from './SoldSneakers'

function App() {

  return (
    <div>
      <div style={{display:"flex", justifyContent:"space-between"}}>
        <div><Link to="/">Sneaker Inventory</Link></div>
        <div><Link to="/add">Add Sneaker</Link></div>
        <div><Link to="/sell">Sell Sneaker</Link></div>
        <div>
          <Link to="/sold">Sold Sneakers</Link>
        </div>
      </div>
      <hr />
      <Routes>
        <Route path="/" element={<Sneakers />} />
        <Route path="/add" element={<AddSneaker />} />
        <Route path="/sell" element={<SellSneaker />} />
        <Route path="/sold" element={<SoldSneakers />} />
      </Routes>
    </div>
  )
}

export default App
