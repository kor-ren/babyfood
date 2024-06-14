import { Outlet } from 'react-router-dom'
import Header from './Header'
import './root.css'

export default function Root() {
    return (
        <>
            <div className="root">
                <Header />
                <div className="content">
                    <Outlet />
                </div>
            </div>
        </>
    )
}