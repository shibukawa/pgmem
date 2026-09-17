package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pgss_shmem_shutdown(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v260 int32
	_ = v260
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v3
	if l0 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v11 + int32(32)
	return
L2:
	;
	v15 = int32(0)
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_shmem_shutdown[0]))
	if v17 == v15 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_shmem_shutdown[1]))
	if v21 == int32(0) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pgss_shmem_shutdown[2])))
	if v25&int32(1) == int32(0) {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v32 = F_AllocateFile(m, int32(_a_F_pgss_shmem_shutdown_0), int32(_a_F_pgss_shmem_shutdown_1))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v260 = F_unlink(m, int32(_a_F_pgss_shmem_shutdown_2))
	mBase = m.M
	goto L1
L7:
	;
	v232 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L8
	} else {
		goto L50
	}
L8:
	;
	return
L9:
	;
	if v32 == int32(0) {
		v222 = v15
		v225 = v3
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v39 = F_fwrite(m, int32(_a_F_pgss_shmem_shutdown_3), int32(4), int32(1), v32)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L8
	} else {
		goto L13
	}
L11:
	;
	v222 = v32
	v225 = v217
	goto L7
L12:
	;
	v217 = int32(0)
	goto L11
L13:
	;
	if v39 != int32(1) {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v46 = F_fwrite(m, int32(_a_F_pgss_shmem_shutdown_4), int32(4), int32(1), v32)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L8
	} else {
		goto L15
	}
L15:
	;
	if v46 != int32(1) {
		goto L12
	} else {
		goto L16
	}
L16:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_shmem_shutdown[1]))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v53)+412))
	if v55 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v118
	v120 = int32(4)
	v124 = F_fwrite(m, v11+v120, v120, int32(1), v32)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L8
	} else {
		goto L21
	}
L18:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v53)+376))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v53)+364))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v53)+352))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v53)+340))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v53)+328))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v53)+316))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v53)+304))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v53)+292))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v53)+280))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v53)+268))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v53)+256))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v53)+244))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v53)+232))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v53)+220))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v53)+208))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v53)+196))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v53)+184))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v53)+172))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v53)+160))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v53)+148))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v53)+136))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v53)+124))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v53)+112))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v53)+100))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v53)+88))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v53)+76))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v53)+64))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v53)+52))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v53)+40))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v53)+28))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v53)+16))
	v118 = v56 + (v57 + (v58 + (v59 + (v60 + (v61 + (v62 + (v63 + (v64 + (v65 + (v66 + (v67 + (v68 + (v69 + (v70 + (v71 + (v72 + (v73 + (v74 + (v75 + (v76 + (v77 + (v78 + (v79 + (v80 + (v81 + (v82 + (v83 + (v84 + (v85 + (v86 + v54))))))))))))))))))))))))))))))
	goto L20
L19:
	;
	v118 = v54
	goto L20
L20:
	;
	goto L17
L21:
	;
	if v124 != int32(1) {
		goto L12
	} else {
		goto L22
	}
L22:
	;
	v130 = F_qtext_load_file(m, v11+int32(28))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L8
	} else {
		goto L23
	}
L23:
	;
	if v130 == int32(0) {
		goto L12
	} else {
		goto L24
	}
L24:
	;
	v135 = v11 + int32(8)
	v137 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_shmem_shutdown[1]))
	F_hash_seq_init(m, v135, v137)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L8
	} else {
		goto L25
	}
L25:
	;
	v140 = F_hash_seq_search(m, v135)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L8
	} else {
		goto L26
	}
L26:
	;
	if v140 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	v143 = v140
	goto L30
L28:
	;
	goto L29
L29:
	;
	v193 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_shmem_shutdown[0]))
	v198 = F_fwrite(m, v193+int32(40), int32(16), int32(1), v32)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L8
	} else {
		goto L45
	}
L30:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v143)+396))
	if v151 < int32(0) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	goto L29
L32:
	;
	v182 = F_hash_seq_search(m, v11+int32(8))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L8
	} else {
		goto L43
	}
L33:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v143)+392))
	v155 = v154 + v151
	if base.Ui32(v142) <= base.Ui32(v155) {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130+v155))))
	if v158 != 0 {
		goto L32
	} else {
		goto L35
	}
L35:
	;
	v161 = F_fwrite(m, v143, int32(432), int32(1), v32)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L8
	} else {
		goto L36
	}
L36:
	;
	if v161 == int32(1) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v166 = int32(1)
	v168 = v151 + v166
	v169 = F_fwrite(m, v130+v154, v166, v168, v32)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L8
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	F_hash_seq_term(m, v11+int32(8))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L8
	} else {
		goto L42
	}
L40:
	;
	if v169 == v168 {
		goto L32
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v217 = v130
	goto L11
L43:
	;
	if v182 != 0 {
		v143 = v182
		goto L30
	} else {
		goto L44
	}
L44:
	;
	goto L31
L45:
	;
	if v198 != int32(1) {
		v217 = v130
		goto L11
	} else {
		goto L46
	}
L46:
	;
	F_emscripten_builtin_free(m, v130)
	mBase = m.M
	v203 = int32(0)
	v205 = F_FreeFile(m, v32)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L8
	} else {
		goto L47
	}
L47:
	;
	if v205 != 0 {
		v222 = v203
		v225 = v203
		goto L7
	} else {
		goto L48
	}
L48:
	;
	v210 = F_durable_rename(m, int32(_a_F_pgss_shmem_shutdown_0), int32(_a_F_pgss_shmem_shutdown_5), int32(15))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L8
	} else {
		goto L49
	}
L49:
	;
	goto L6
L50:
	;
	if v232 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L8
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	F_emscripten_builtin_free(m, v225)
	mBase = m.M
	if v222 != 0 {
		goto L57
	} else {
		goto L58
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(_a_F_pgss_shmem_shutdown_0)
	F_errmsg(m, int32(_a_F_pgss_shmem_shutdown_6), v11)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L8
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(_a_F_pgss_shmem_shutdown_7), int32(824), int32(_a_F_pgss_shmem_shutdown_8))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L8
	} else {
		goto L56
	}
L56:
	;
	goto L53
L57:
	;
	v247 = F_FreeFile(m, v222)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L8
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v250 = F_unlink(m, int32(_a_F_pgss_shmem_shutdown_0))
	mBase = m.M
	goto L6
L60:
	;
	goto L59
}
