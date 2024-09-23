// https://qiita.com/mu_tomoya/items/7545bea039e82e483f9e

// // // Filename - index.js
import React, { useState, useEffect } from 'react';
import Calendar from 'react-calendar';
import 'react-calendar/dist/Calendar.css';
import EventForm from '@/components/EventForm';
import axios from 'axios';

interface Event {
    id: number;
    title: string;
    date: string;
    hour: string;
    user: string; // ユーザー名の追加
}


const generateHourList = () => {
    const hourList = [];
    for (let hour = 0; hour < 24; hour++) {
        const formattedHour = String(hour).padStart(2, '0');
        hourList.push(`${formattedHour}:00`);
    }
    return hourList;
};

export default function CalendarGfg() {
    const [value, onChange] = useState<Date>(new Date());
    const [events, setEvents] = useState<Event[]>([]);
    const [selectedDate, setSelectedDate] = useState<Date | null>(null);
    const [editingEvent, setEditingEvent] = useState<Event | null>(null);

    const handleClickDay = (value: Date) => {
        setSelectedDate(value);
        setEditingEvent(null);
    };

    const formatDate = (date: Date | null) => {
        if (!date) {
            return '';
        }
        const year = date.getFullYear();
        const month = String(date.getMonth() + 1).padStart(2, '0');
        const day = String(date.getDate()).padStart(2, '0');
        return `${year}-${month}-${day}`;
    };

    const tileContent = ({ date, view }: { date: Date; view: string }) => {
        if (view === 'month') {
            const formattedDate = formatDate(date);
            const eventCount = events.filter(event => event.date === formattedDate).length;
            return eventCount > 0 ? <div>{eventCount}</div> : null;
        }
    };

    const fetchEvents = async () => {
        try {
            const response = await axios.get('/events');
            setEvents(response.data);
        } catch (error) {
            console.error('Error fetching events:', error);
        }
    };

    const addEvent = async (event: Event) => {
        try {
            await axios.post('/events', event);
            fetchEvents(); // イベントデータを再取得
        } catch (error) {
            console.error('Error adding event:', error);
        }
    };

    const editEvent = async (event: Event) => {
        try {
            await axios.put(`/events/${event.id}`, event);
            fetchEvents(); // イベントデータを再取得
            setEditingEvent(null);
        } catch (error) {
            console.error('Error editing event:', error);
        }
    };

    const deleteEvent = async (eventId: number) => {
        try {
            await axios.delete(`/events/${eventId}`);
            fetchEvents(); // イベントデータを再取得
        } catch (error) {
            console.error('Error deleting event:', error);
        }
    };

    useEffect(() => {
        fetchEvents();
    }, []);

    return (
        <div className="flex">
            <div className="w-1/3 m-4">
                <Calendar
                    onChange={onChange}
                    value={value}
                    locale="ja-JP"
                    onClickDay={handleClickDay}
                    tileContent={tileContent}
                />
            </div>
            <div className="w-2/3 m-4">
                {selectedDate && (
                    <div className="bg-white shadow-md rounded-md p-4">
                        <h2 className="text-lg font-bold mb-4">{formatDate(selectedDate)}</h2>
                        <ul>
                            {generateHourList().map((hour) => (
                                <li key={hour} className="flex items-center justify-between mb-2">
                                    <span className="font-semibold">{hour}:</span>
                                    <div className="flex-grow flex items-center justify-center">
                                        {events.find(
                                            (event) => event.date === formatDate(selectedDate) && event.hour === hour
                                        ) && (
                                            <div>
                                                {events.find(
                                                    (event) => event.date === formatDate(selectedDate) && event.hour === hour
                                                )?.title}{' '}
                                                ({events.find(
                                                    (event) => event.date === formatDate(selectedDate) && event.hour === hour
                                                )?.user || '名無し'})
                                            </div>
                                        )}
                                        {!events.find(
                                            (event) => event.date === formatDate(selectedDate) && event.hour === hour
                                        ) && (
                                            <div className="flex items-center justify-center">
                                                <span className="text-gray-500">予定なし</span>
                                                <EventForm
                                                    onAddEvent={addEvent}
                                                    selectedHour={hour}
                                                    selectedDate={selectedDate}
                                                />
                                            </div>
                                        )}
                                        {events.find(
                                            (event) => event.date === formatDate(selectedDate) && event.hour === hour
                                        ) && (
                                            <>
                                                <button
                                                    className="ml-2 bg-yellow-500 hover:bg-yellow-700 text-white font-semibold py-1 px-2 rounded"
                                                    onClick={() =>
                                                        setEditingEvent(
                                                            events.find(
                                                                (event) => event.date === formatDate(selectedDate) && event.hour === hour
                                                            )!
                                                        )
                                                    }
                                                >
                                                    編集
                                                </button>
                                                <button
                                                    className="ml-2 bg-red-500 hover:bg-red-700 text-white font-semibold py-1 px-2 rounded"
                                                    onClick={() =>
                                                        deleteEvent(
                                                            events.find(
                                                                (event) => event.date === formatDate(selectedDate) && event.hour === hour
                                                            )!.id
                                                        )
                                                    }
                                                >
                                                    削除
                                                </button>
                                            </>
                                        )}
                                        {editingEvent?.date === formatDate(selectedDate) && editingEvent?.hour === hour && (
                                            <EventForm
                                                onEditEvent={editEvent}
                                                selectedHour={hour}
                                                selectedDate={selectedDate}
                                                editMode
                                                initialTitle={editingEvent.title}
                                                initialUser={editingEvent.user}
                                            />
                                        )}
                                    </div>
                                </li>
                            ))}
                        </ul>
                    </div>
                )}
            </div>
        </div>
    );
}