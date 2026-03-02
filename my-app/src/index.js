import { createRoot } from 'react-dom/client';
import LikeCounter from './components/LikeCounter'; // 👈 파일 가져오기 (확장자 .js는 생략 가능)
import WindowWidthCounter from './components/WindowWitdhCounter';
const items = [
    { id: 1, name: "apple", stock: 5 },
    { id: 2, name: "grape", stock: 3 },
    { id: 3, name: "watermelom", stock: 0}
]

const root = createRoot(document.getElementById('root'));
root.render(
    <>
    <div>
        {items.map((item) => (
            <LikeCounter
                key={item.id}
                title={item.name}
                initCount={item.stock}
            />
            ))}
    </div>
    <div>
        <WindowWidthCounter />
    </div>
    </>
);