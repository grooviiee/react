import { useState, useEffect } from 'react';

function WindowWidthCounter() {
    const [width, setWidth] = useState(window.innerWidth);

    useEffect(() => {
        const handleResize = () => {
            setWidth(window.innerWidth);
            console.log("Window size resized.");
        };

        window.addEventListener('resize', handleResize);
        return () => {
            window.removeEventListener('resize', handleResize);
        };
    }, []);

    return (
        <div>
            <p>현재 창 너비: {width}px</p>
            {width < 600 && <p style={{ color: 'blue' }}>📱 모바일 화면 모드입니다.</p>}
        </div>
    );
}

export default WindowWidthCounter; // 👈 다른 파일에서 쓸 수 있게 내보내기