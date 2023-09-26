import { useCallback, useState } from 'react';

function useCheckRoom() {
  const [isValid, setIsValid] = useState<boolean>(false);

  const checkRoom = useCallback(async (roomName: string, config?: string) => {
    try {
      if (typeof roomName !== 'string' || !roomName) return false;
      const result = await fetch(`https://faceto-ai.withcontext.ai/check/${roomName}`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: config ? JSON.stringify({ c: config }) : '',
      }).then((res) => res.json());
      setIsValid(result.valid);
      return result.valid as boolean;
    } catch (e) {
      setIsValid(false);
      return false;
    }
  }, []);

  return {
    isValid,
    checkRoom,
  };
}

export default useCheckRoom;
