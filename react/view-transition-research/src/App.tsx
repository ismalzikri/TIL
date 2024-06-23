import { useState } from "react";
import { flushSync } from "react-dom";
import "./App.css";

function App() {
  const [items, setItems] = useState(() => {
    let items = [];
    for (let i = 0; i < 30; i++) {
      items.push({ id: i, name: "Item " + i });
    }
    return items;
  });

  let add = () => {
    document.startViewTransition(() => {
      flushSync(() => {
        setItems((items) => [{ id: Date.now(), name: "New item" }, ...items]);
      });
    });
  };

  let remove = (i) => {
    document.startViewTransition(() => {
      flushSync(() => {
        setItems((items) => [...items.slice(0, i), ...items.slice(i + 1)]);
      });
    });
  };

  return (
    <>
      <button onClick={add}>+</button>
      <ul>
        {items.map((item, i) => (
          <li key={item.id} style={{ viewTransitionName: `item-${item.id}` }}>
            {item.name}
            <button onClick={() => remove(i)}>x</button>
          </li>
        ))}
      </ul>
    </>
  );
}

export default App;
