import { useState } from "react";
import { flushSync } from "react-dom";
import "./App.css";

type Item = {
  id: number;
  name: string;
};

function App() {
  const [items, setItems] = useState<Item[]>(() => {
    const initialItems: Item[] = [];
    for (let i = 1; i <= 30; i++) {
      initialItems.push({ id: i, name: "Item " + i });
    }
    return initialItems;
  });

  const add = () => {
    document.startViewTransition(() => {
      flushSync(() => {
        setItems((items) => [{ id: Date.now(), name: "New item" }, ...items]);
      });
    });
  };

  const remove = (index: number) => {
    document.startViewTransition(() => {
      flushSync(() => {
        setItems((items) => [
          ...items.slice(0, index),
          ...items.slice(index + 1),
        ]);
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
