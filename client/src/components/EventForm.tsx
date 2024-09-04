// components/EventForm.jsx
import React, { useState } from 'react';

interface EventFormProps {
    onAddEvent: (event: Event) => void;
    onEditEvent?: (event: Event) => void;
    selectedHour: string;
    selectedDate: Date;
    editMode?: boolean;
    initialTitle?: string;
    initialUser?: string; // 初期ユーザー名の追加
    disabled?: boolean;
}

const EventForm: React.FC<EventFormProps> = ({
    onAddEvent,
    onEditEvent,
    selectedHour,
    selectedDate,
    editMode = false,
    initialTitle = '',
    initialUser = '', // 初期ユーザー名のデフォルト値は空文字列
    disabled = false,
}) => {
    const [eventTitle, setEventTitle] = useState(editMode ? initialTitle : '');
    const [eventUser, setEventUser] = useState(editMode ? initialUser : ''); // ユーザー名のstate追加

    const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        const event = { title: eventTitle, date: formatDate(selectedDate), hour: selectedHour, user: eventUser }; // ユーザー名を追加
        if (editMode && onEditEvent) {
            onEditEvent(event);
        } else {
            onAddEvent(event);
        }
        setEventTitle('');
        setEventUser(''); // フォームをリセット
    };

    const formatDate = (date: Date) => {
        const year = date.getFullYear();
        const month = String(date.getMonth() + 1).padStart(2, '0');
        const day = String(date.getDate()).padStart(2, '0');
        return `${year}-${month}-${day}`;
    };

    return (
        <form onSubmit={handleSubmit} className="flex items-center ml-2">
            <select
                value={eventTitle}
                onChange={(e) => setEventTitle(e.target.value)}
                className="py-1 px-2 border rounded-l"
                disabled={disabled}
            >
                <option value="">予定を選択</option>
                <option value="会議">会議</option>
                <option value="打ち合わせ">打ち合わせ</option>
                <option value="休憩">休憩</option>
            </select>
            <input
                type="text"
                placeholder="ユーザー名"
                value={eventUser}
                onChange={(e) => setEventUser(e.target.value)}
                className="py-1 px-2 border"
                disabled={disabled}
            />
            <button
                type="submit"
                className={`py-1 px-2 rounded-r ${
                    disabled || eventTitle === '' || eventUser === ''
                        ? 'bg-gray-400 cursor-not-allowed'
                        : 'bg-blue-500 hover:bg-blue-700 text-white font-semibold'
                }`}
                disabled={disabled || eventTitle === '' || eventUser === ''}
            >
                {editMode ? '更新' : '追加'}
            </button>
        </form>
    );
};

export default EventForm;