import matplotlib.pyplot as plt
import numpy as np
import pandas as pd
from os import path
from os import listdir
import re

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
        data_path = ''
        
        if file.endswith('.csv'):
            data_path = path.join('..', 'results', file)
        else:
            data_path = path.join('..', 'results', f'{file}.csv')
        
        df = pd.read_csv(data_path)
        draw_plot(df, file,  x_label, y_label, e_label)

        data_file_name = extract_data_file_name(file)

        if data_file_name != '':
            data_file_name += '_'

    graph_path = path.join('..', 'graphs', f'{data_file_name}{x_label}_{y_label}.png')
    plt.savefig(graph_path)

def draw_for_prefixed_files(files_name_prefix, x_label, y_label):
    only_files = [f for f in listdir(path.join('..', 'results')) if path.isfile(path.join(path.join('..', 'results'), f))]
    filtered_files = filter(lambda x: x.startswith(files_name_prefix), only_files)
    draw_for_files(filtered_files, x_label, y_label)

def extract_data_file_name(fileName):
    return re.findall(r'(?<=-)[^.\s]*|$', fileName)[0]

if __name__ == '__main__':
    swap_gs_files = ['swapGreedy', 'swapSteepest']
    step_processing_prefix = "step"
    similiarity_files_prefix = "similarity"
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

    draw_for_prefixed_files(step_processing_prefix, 'iteration_num', 'quality')
    draw_for_prefixed_files(similiarity_files_prefix, 'quality', 'similarity')
