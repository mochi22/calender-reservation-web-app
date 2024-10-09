// https://qiita.com/mu_tomoya/items/7545bea039e82e483f9e

// // // Filename - index.js
import React, { useState, useEffect } from 'react';
import Calendar from 'react-calendar';
import 'react-calendar/dist/Calendar.css';
import EventForm from '@/components/EventForm';
import axios from 'axios';

interface EventData {
    id: string;
    title: string;
    user: string;
    date: string;
    hour: string;
    createat: string;
    updateat: string;
}

// define hour list from 0 to 24 every 1 hour
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
    const [events, setEvents] = useState<EventData[]>([]);
    const [selectedDate, setSelectedDate] = useState<Date | null>(null);
    const [editingEvent, setEditingEvent] = useState<EventData | null>(null);

    // handle click day in calendar
    const handleClickDay = (value: Date) => {
        setSelectedDate(value);
        console.log("setSelectedDate value:", value);
        setEditingEvent(null);
    };

    // date formatted
    const formatDate = (date: Date | null) => {
        if (!date) {
            return '';
        }
        const year = date.getFullYear();
        const month = String(date.getMonth() + 1).padStart(2, '0');
        const day = String(date.getDate()).padStart(2, '0');
        return `${year}-${month}-${day}`;
    };

    // define tile param in react-calendar
    const tileContent = ({ date, view }: { date: Date; view: string }) => {
        if (view === 'month') {
            const formattedDate = formatDate(date);
            const eventCount = events.filter(event => event.date === formattedDate).length;
            return eventCount > 0 ? <div>{eventCount}</div> : null;
        }
    };

    // define backend port
    const root_path = "http://localhost:8080";

    const fetchEvents = async () => {
        try {
            // get all data from db!
            const response = await axios.get(`${root_path}/events`);
            if (response.data != null) {
                setEvents(response.data); //store all data
                console.log("res:", response.data);
            }
        } catch (error) {
            console.error('Error fetching events:', error);
        }
    };

    const addEvent = async (event: EventData) => {
        try {
            //adding reserved data to db
            await axios.post(`${root_path}/events`, event);
            fetchEvents(); // イベントデータを再取得
        } catch (error) {
            console.error('Error adding event:', error);
        }
    };

    const editEvent = async (event: EventData) => {
        try {
            await axios.put(`${root_path}/events/${event.id}`, event);
            fetchEvents(); // イベントデータを再取得
            setEditingEvent(null);
        } catch (error) {
            console.error('Error editing event:', error);
        }
    };

    const deleteEvent = async (eventId: number) => {
        try {
            await axios.delete(`${root_path}/events/${eventId}`);
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
                    <div className="bg-green shadow-md rounded-md p-4">
                        <h2 className="text-lg font-bold mb-4">{formatDate(selectedDate)}</h2>
                        <ul>
                            {generateHourList().map((hour) => (
                                <li key={hour} className="flex items-center justify-between mb-2">
                                    <span className="font-semibold">{hour}:</span>
                                    <div className="flex-grow flex items-center justify-center">
                                        {
                                            events.find(
                                                (event) => event.date === formatDate(selectedDate) && event.hour === hour
                                            ) && (
                                                <div>
                                                    {
                                                        events.find(
                                                            (event) => event.date === formatDate(selectedDate) && event.hour === hour
                                                        )?.title
                                                    }{' '}
                                                    ({
                                                        events.find(
                                                            (event) => event.date === formatDate(selectedDate) && event.hour === hour
                                                        )?.username || '名無し'
                                                    })
                                                </div>
                                            )
                                        }
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