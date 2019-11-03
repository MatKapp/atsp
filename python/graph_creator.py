import matplotlib.pyplot as plt
import numpy as np
import pandas as pd
from os import path


def draw_plot(df, label, x_label, y_label, e_label=None):
    # Plotting the data
    x = df[x_label]
    y = df[y_label]

    if e_label:
        e = df[e_label]
        plt.errorbar(x, y, e, label=label, marker='^')
    else:
        plt.plot(x, y, label=label, marker='^')

    # Adding a legend
    plt.legend()


def draw_for_files(files, x_label, y_label, e_label=None):
    plt.clf()
    for file in files:
        data_path = path.join('..', 'results', f'{file}.csv')
        df = pd.read_csv(data_path)
        draw_plot(df, file,  x_label, y_label, e_label)
    graph_path = path.join('..', 'graphs', f'{x_label}_{y_label}.png')
    plt.savefig(graph_path)


if __name__ == '__main__':
    swap_gs_files = ['swapGreedy', 'swapSteepest']
    all_gs_files = swap_gs_files + ['reverseGreedy', 'reverseSteepest']
    gsr_files = swap_gs_files + ['random']
    gsrh_files = gsr_files + ['heuristic']
    draw_for_files(gsrh_files, 'size', 'best')
    draw_for_files(all_gs_files, 'size', 'best')
    draw_for_files(gsr_files, 'size', 'time')
    draw_for_files(swap_gs_files, 'size', 'mean', 'std')
    draw_for_files(swap_gs_files, 'size', 'mean_steps')
    draw_for_files(all_gs_files, 'size', 'mean_steps')
    draw_for_files(all_gs_files, 'size', 'reviewed_solutions')
    draw_for_files(all_gs_files, 'size', 'quality_time')
