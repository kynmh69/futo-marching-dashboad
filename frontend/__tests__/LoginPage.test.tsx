import React from 'react';
import { render, screen } from '@testing-library/react';
import LoginPage from '../app/login/page';

// Mock the useRouter hook
jest.mock('next/navigation', () => ({
  useRouter: () => ({
    push: jest.fn(),
  }),
}));

// Mock the useAuth hook
jest.mock('../lib/auth', () => ({
  useAuth: () => ({
    login: jest.fn(),
  }),
}));

describe('LoginPage', () => {
  it('renders the login form', () => {
    render(<LoginPage />);
    
    // Check if important elements are rendered
    expect(screen.getByText('FUTO Marching Dashboard')).toBeInTheDocument();
    expect(screen.getByText('Sign in to your account')).toBeInTheDocument();
    expect(screen.getByLabelText('Username')).toBeInTheDocument();
    expect(screen.getByLabelText('Password')).toBeInTheDocument();
    expect(screen.getByRole('button', { name: 'Sign in' })).toBeInTheDocument();
  });
});