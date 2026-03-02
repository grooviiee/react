import { useState, useEffect } from 'react';

function LikeCounter(props) {
    // 0 대신 props.initCount를 초기값으로 설정합니다.
    const [count, setCount] = useState(props.initCount);

    useEffect(() => {
        console.log(`${props.title} Component appear!`);
    }, [])

    const increase = () => setCount(count + 1);
    const decrease = () => {
        if (count > 0) setCount(count - 1);
    };



    return (
        <div>
            <h2>{props.title}</h2>
            {/* 조건부 렌더링 적용! */}
            {count === 0 ? <p style={{color: 'red'}}>품절되었습니다</p> : <p>남은 수량: {count}</p>}
            
            <button onClick={() => setCount(count + 1)}>➕</button>
            {/* 0보다 클 때만 감소 버튼 보여주기 */}
            {count > 0 && <button onClick={() => setCount(count - 1)}>➖</button>}
        </div>
    );
}

export default LikeCounter; // 👈 다른 파일에서 쓸 수 있게 내보내기