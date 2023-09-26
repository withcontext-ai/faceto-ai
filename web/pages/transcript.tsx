import { format, isToday, isYesterday } from 'date-fns';
import { Avatar, AvatarFallback, AvatarImage } from '../components/ui/avatar';
import { cn, getAvatarBgColor } from '../lib/utils';
import { useEffect, useMemo, useState } from 'react';
import { Button } from '../components/ui/button';

const formatTime = (time: number) => {
  if (isToday(time)) {
    return format(time, 'kk:mm:ss aa');
  }
  if (isYesterday(time)) {
    return `Yesterday at ${format(time, 'kk:mm:ss aa')}`;
  } else return format(time, 'dd/MM/yyyy kk:mm:ss aa');
};

const botAvatar = '/assets/bot_avatar.png';

const Item = ({ is_bot, name, text, timestamp }: any) => {
  const isUser = !is_bot;

  const color = useMemo(() => getAvatarBgColor(name), [name]);


  return (
    <div className="flex flex-col">
      <div className="flex space-x-4">
        <div className="relative shrink-0">
          <Avatar className="relative h-12 w-12">
            <AvatarImage src={isUser ? '' : botAvatar} alt={name} />
            <AvatarFallback className={`bg-${color}-600`}>{name}</AvatarFallback>
          </Avatar>
        </div>

        <div className="flex flex-col">
          <div className="mb-5 flex items-center space-x-2">
            <div className="text-white text-sm">{name}</div>
            <div className="text-gray-500 text-xs">{formatTime(timestamp * 1000)}</div>
          </div>
          <div className="flex items-end">
            <div
              className={cn(
                'rounded-lg p-4 text-sm sm:max-w-full md:max-w-full lg:max-w-3xl	xl:max-w-3xl',
                isUser ? 'bg-gray-100 text-black' : 'bg-[#0F3FEF]',
              )}
            >
              {text}
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

const List = ({ items }: any) => {
  if (items.length === 0) {
    return <div className="text-white text-center">No transcript available</div>;
  }

  return items.map(({ is_bot, name, text, timestamp }: any) => (
    <Item key={timestamp} is_bot={is_bot} name={name} text={text} timestamp={timestamp} />
  ));
};

const DownloadButton = ({ items }: { items: any }) => {
  const text = useMemo(() => {
    return items
      .map(({ name, text, timestamp }: any) => `${name} ${formatTime(timestamp * 1000)}: ${text}`)
      .join('\n');
  }, [items]);

  function handleSubmit(e: any) {
    e.preventDefault();

    var element = document.createElement('a');
    element.setAttribute('href', 'data:text/plain;charset=utf-8,' + encodeURIComponent(text));
    element.setAttribute('download', 'transcript.txt');

    element.style.display = 'none';
    document.body.appendChild(element);

    element.click();

    document.body.removeChild(element);
  }

  if (!text) return null;

  return (
    <form onSubmit={handleSubmit}>
      <Button type="submit" variant="outline" size="sm">
        Download
      </Button>
    </form>
  );
};

function useTranscript() {
  const [items, setItems] = useState([]);

  useEffect(() => {
    const str = sessionStorage.getItem('transcript');
    if (str) {
      try {
        const json = JSON.parse(str);
        setItems(json);
      } catch (error) {
        console.log(error);
      }
    }
  }, []);

  return items;
}

function TranscriptPage() {
  const items = useTranscript();

  return (
    <div className="bg-[#282828] min-h-screen flex">
      <div className="bg-black max-w-[600px] mx-auto p-6 flex-1">
        <div className="flex justify-between items-center">
          <h1 className="text-white text-lg font-semibold">Transcript</h1>
          <DownloadButton items={items} />
        </div>
        <div className="space-y-6 mt-12">
          <List items={items} />
        </div>
      </div>
    </div>
  );
}

export default TranscriptPage;
