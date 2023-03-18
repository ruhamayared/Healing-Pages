package main

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func CreateEntry(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		entry := new(Entry)
		if err := c.Bind(entry); err != nil {
			return c.String(http.StatusBadRequest, "invalid entry")
		}
		entry.CreatedAt = time.Now()

		// Insert entry into database
		if _, err := db.Exec("INSERT INTO entries (entry, created_at) VALUES ($1, $2)", entry.Entry, entry.CreatedAt); err != nil {
			return err
		}

		return c.JSON(http.StatusCreated, entry)
	}
}

func GetEntry(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "invalid id")
		}

		// Get entry from database
		row := db.QueryRow("SELECT id, entry, created_at FROM entries WHERE id=$1", id)
		entry := new(Entry)
		if err := row.Scan(&entry.ID, &entry.Entry, &entry.CreatedAt); err != nil {
			if err == sql.ErrNoRows {
				return c.String(http.StatusNotFound, "entry not found")
			}
			return err
		}

		return c.JSON(http.StatusOK, entry)
	}
}

func UpdateEntry(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "invalid id")
		}

		entry := new(Entry)
		if err := c.Bind(entry); err != nil {
			return c.String(http.StatusBadRequest, "invalid entry")
		}
		entry.ID = id

		// Update entry in database
		if _, err := db.Exec("UPDATE entries SET entry=$1 WHERE id=$2", entry.Entry, entry.ID); err != nil {
			if err == sql.ErrNoRows {
				return c.String(http.StatusNotFound, "entry not found")
			}
			return err
		}

		return c.JSON(http.StatusOK, entry)
	}
}

func DeleteEntry(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "invalid id")
		}

		// Delete entry from database
		if _, err := db.Exec("DELETE FROM entries WHERE id=$1", id); err != nil {
			if err == sql.ErrNoRows {
				return c.String(http.StatusNotFound, "entry not found")
			}
			return err
		}

		return c.NoContent(http.StatusNoContent)
	}
}

func GetAllEntries(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		rows, err := db.Query("SELECT id, entry, created_at FROM entries")
		if err != nil {
			return err
		}
		defer rows.Close()

		entries := make([]*Entry, 0)
		for rows.Next() {
			entry := new(Entry)
			if err := rows.Scan(&entry.ID, &entry.Entry, &entry.CreatedAt); err != nil {
				return err
			}
			entries = append(entries, entry)
		}

		return c.JSON(http.StatusOK, entries)
	}
}
