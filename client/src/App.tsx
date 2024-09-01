import React, {useState, useEffect} from 'react';
import './App.css';
import axios from "axios";

type User = {
  id:number;
	name: string;
	password: string;
	calender_id: number;
}
// var users = [
// 	{id: 1, name: "apple", password: "pass1", calender_id: 10},
// 	{id: 2, name: "banana", password: "pass22", calender_id: 20},
// 	{id: 3, name: "grape", password: "pass333", calender_id: 10},
// ]

function App() {
  const [users, setUsers] = useState<User []>([{id: 0, name: "", password: "", calender_id: 0}])
  useEffect(() => {
    (
      async () => {
        const data = await axios.get("http://localhost:8080")
        console.log("kuraryu logging")
        console.log(data)
        console.log(data.data)
        setUsers(data.data)
      }
    )()
  }, [])

  return (
    <div>
      {users.map(user => (
        <p key={user.id}>
          <span>/user_name:{user.name}</span>
          <span>/user_password:{user.password}</span>
          <span>/user_calender_id:{user.calender_id}</span>
        </p>
      ))}
    </div>
  );
}

export default App;
