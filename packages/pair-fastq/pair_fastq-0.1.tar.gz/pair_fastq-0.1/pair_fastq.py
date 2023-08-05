import click, json, pandas as pd, os, numpy as np



#######################
#### RUN CLI GROUP ####

CONTEXT_SETTINGS = dict(help_option_names=['-h', '--help'], max_content_width=90)

@click.group(context_settings=CONTEXT_SETTINGS)
@click.version_option(version='0.1')
def pairfastq():
    pass


##########################
#### PAIR FASTQ FILES ####

@pairfastq.command()
@click.option('-fq','--fastq',help='Fastq file(s) for pairing. Comma delimited.', type=click.Path(exists=False))
@click.option('-csv','--config',default = './config.csv', show_default=True, help='Csv configuration file specifying reads 1 data, reads 2 data, and conditions data in columns r1, r2, and condition, respectively.',type=click.Path(exists=False))
@click.option('-o','--output_file',default = 'cwl.output.json', show_default=True, help='Cwl output json file array, essentially a json containing fastq file pairs.') # remember to inherit metadata from the input files
@click.option('-w','--work_dir',default = './', show_default=True, help='Work directory.',type=click.Path(exists=False))
def pair_fastq(fastq,config,output_file,work_dir):
    fastq = fastq.split(',')
    df = pd.read_csv(config)
    fastq_pairs = []
    empty_pairs = [['Sample1','Sample2']]
    fastq_pairs_all = df[['r1','r2']].as_matrix().tolist()
    fastq_check = []
    for r1,r2 in fastq_pairs_all:
        if r1 in fastq and r2 in fastq:
            fastq_pairs.append(map(lambda path: {'class':'File','path':os.path.abspath(path)},[r1,r2]))
            fastq_check.append(r1)
        else:
            empty_pairs.append([r1,r2])

    empty_pairs.extend([(fq,'') for fq in set(fastq)-set(reduce(lambda x, y: x+y, fastq_pairs_all))])

    empty_pairs = '\n'.join('\t'.join(empty_pair) for empty_pair in empty_pairs)

    with open(work_dir+'empty_pairs.txt','w') as f:
        f.write(empty_pairs)

    fastq_pairs_dict = {"fastq_pairs":fastq_pairs,"empty_pairs":{"class":'File',"path":os.path.abspath(work_dir+'/')+'empty_pairs.txt'}}

    df = df[df['r1'].isin(fastq_check).as_matrix()]
    if 'condition' in list(df):
        df = df.iloc[np.argsort(df['condition']),:].reset_index(drop=True)
    df['r1','r2'] = df[['r1','r2']].map(os.path.abspath)
    df.to_csv(work_dir+'output.paired_fastq.csv')

    with open(output_file,'w') as f:
        json.dump(fastq_pairs_dict,f)


#### RUN CLI ####

if __name__ == '__main__':
    pairfastq()