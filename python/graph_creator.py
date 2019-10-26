import matplotlib.pyplot as plt
import numpy as np
import pandas as pd


def draw_plot(df, x_label, y_label, label):
    # Plotting the data
    x = df[x_label]
    y = df[y_label]
    plt.plot(x, y, label=label)

    # Adding a legend
    plt.legend()


def draw_for_files(files, x_label, y_label):
    for file in files:
        df = pd.read_csv(f'../results/{file}.csv')
        draw_plot(df, x_label, y_label, file)
    plt.show()


if __name__ == '__main__':
    gs_files = ['greedy', 'steepest']
    gsr_files = gs_files + ['random']
    gsrh_files = gsr_files + ['heuristic']
    draw_for_files(gsrh_files, 'size', 'best')
    draw_for_files(gsr_files, 'size', 'time')
    draw_for_files(gs_files, 'size', 'mean')
