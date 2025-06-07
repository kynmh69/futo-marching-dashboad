"use client";

import { useAuth } from "@/lib/auth";

export default function DashboardPage() {
  const { user } = useAuth();

  return (
    <div>
      <h1 className="text-2xl font-bold mb-6">Dashboard</h1>
      
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <div className="bg-white p-6 rounded-lg shadow">
          <h2 className="text-lg font-semibold mb-2">Welcome Back!</h2>
          <p className="text-gray-600">
            Hello, {user?.fullName || "User"}! This is your dashboard where you can manage your tasks, view the calendar, and access practice menus.
          </p>
        </div>
        
        <div className="bg-white p-6 rounded-lg shadow">
          <h2 className="text-lg font-semibold mb-2">Quick Actions</h2>
          <div className="space-y-2">
            <a href="/dashboard/tasks" className="block p-2 text-indigo-600 hover:bg-indigo-50 rounded">
              View Tasks
            </a>
            <a href="/dashboard/calendar" className="block p-2 text-indigo-600 hover:bg-indigo-50 rounded">
              Check Calendar
            </a>
            <a href="/dashboard/practice" className="block p-2 text-indigo-600 hover:bg-indigo-50 rounded">
              Practice Menu
            </a>
          </div>
        </div>
        
        <div className="bg-white p-6 rounded-lg shadow">
          <h2 className="text-lg font-semibold mb-2">Time Tracking</h2>
          <div className="space-y-4">
            <p className="text-gray-600">Track your practice sessions:</p>
            <div className="flex space-x-4">
              <button className="px-4 py-2 bg-green-500 text-white rounded hover:bg-green-600">
                Clock In
              </button>
              <button className="px-4 py-2 bg-red-500 text-white rounded hover:bg-red-600">
                Clock Out
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}