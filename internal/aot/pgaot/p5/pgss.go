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
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
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
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v253 int32
	_ = v253
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v276 int32
	_ = v276
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
	v18 = *(*int32)(unsafe.Add(mBase, _consts[1491]))
	if v18 == v15 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v23 = *(*int32)(unsafe.Add(mBase, _consts[1492]))
	if v23 == int32(0) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1493])))
	if v27 != int32(1) {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v35 = F_AllocateFile(m, int32(247807), int32(34101))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v276 = F_unlink(m, int32(119666))
	mBase = m.M
	goto L1
L7:
	;
	v242 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L8
	} else {
		goto L50
	}
L8:
	;
	return
L9:
	;
	if v35 == int32(0) {
		v232 = v15
		v235 = v3
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v42 = F_fwrite(m, int32(4120604), int32(4), int32(1), v35)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L8
	} else {
		goto L13
	}
L11:
	;
	v232 = v35
	v235 = v227
	goto L7
L12:
	;
	v227 = int32(0)
	goto L11
L13:
	;
	if v42 != int32(1) {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v49 = F_fwrite(m, int32(4120608), int32(4), int32(1), v35)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L8
	} else {
		goto L15
	}
L15:
	;
	if v49 != int32(1) {
		goto L12
	} else {
		goto L16
	}
L16:
	;
	v54 = *(*int32)(unsafe.Add(mBase, _consts[1492]))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v56)+412))
	if v58 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v123
	v125 = int32(4)
	v129 = F_fwrite(m, v11+v125, v125, int32(1), v35)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L8
	} else {
		goto L21
	}
L18:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v56)+376))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v56)+364))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v56)+352))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v56)+340))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v56)+328))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v56)+316))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v56)+304))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v56)+292))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v56)+280))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v56)+268))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v56)+256))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v56)+244))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v56)+232))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v56)+220))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v56)+208))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v56)+196))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v56)+184))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v56)+172))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v56)+160))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v56)+148))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v56)+136))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v56)+124))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v56)+112))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v56)+100))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v56)+88))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v56)+76))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v56-int32(-64))))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v56)+52))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v56)+40))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v56)+28))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v56)+16))
	v123 = v59 + (v60 + (v61 + (v62 + (v63 + (v64 + (v65 + (v66 + (v67 + (v68 + (v69 + (v70 + (v71 + (v72 + (v73 + (v74 + (v75 + (v76 + (v77 + (v78 + (v79 + (v80 + (v81 + (v82 + (v83 + (v84 + (v87 + (v88 + (v89 + (v90 + (v91 + v57))))))))))))))))))))))))))))))
	goto L20
L19:
	;
	v123 = v57
	goto L20
L20:
	;
	goto L17
L21:
	;
	if v129 != int32(1) {
		goto L12
	} else {
		goto L22
	}
L22:
	;
	v135 = F_qtext_load_file(m, v11+int32(28))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L8
	} else {
		goto L23
	}
L23:
	;
	if v135 == int32(0) {
		goto L12
	} else {
		goto L24
	}
L24:
	;
	v142 = *(*int32)(unsafe.Add(mBase, _consts[1492]))
	F_hash_seq_init(m, v11+int32(8), v142)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L8
	} else {
		goto L25
	}
L25:
	;
	v147 = F_hash_seq_search(m, v11+int32(8))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L8
	} else {
		goto L26
	}
L26:
	;
	if v147 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	v150 = v147
	goto L30
L28:
	;
	goto L29
L29:
	;
	v200 = *(*int32)(unsafe.Add(mBase, _consts[1491]))
	v205 = F_fwrite(m, v200+int32(40), int32(16), int32(1), v35)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L8
	} else {
		goto L45
	}
L30:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v150)+396))
	if v158 < int32(0) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	goto L29
L32:
	;
	v189 = F_hash_seq_search(m, v11+int32(8))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L8
	} else {
		goto L43
	}
L33:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v150)+392))
	v162 = v161 + v158
	if base.Ui32(v149) <= base.Ui32(v162) {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+v162))))
	if v165 != 0 {
		goto L32
	} else {
		goto L35
	}
L35:
	;
	v168 = F_fwrite(m, v150, int32(432), int32(1), v35)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L8
	} else {
		goto L36
	}
L36:
	;
	if v168 == int32(1) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v173 = int32(1)
	v175 = v158 + v173
	v176 = F_fwrite(m, v135+v161, v173, v175, v35)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
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
	v183 = m.ExcPending
	if v183 != 0 {
		goto L8
	} else {
		goto L42
	}
L40:
	;
	if v176 == v175 {
		goto L32
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v227 = v135
	goto L11
L43:
	;
	if v189 != 0 {
		v150 = v189
		goto L30
	} else {
		goto L44
	}
L44:
	;
	goto L31
L45:
	;
	if v205 != int32(1) {
		v227 = v135
		goto L11
	} else {
		goto L46
	}
L46:
	;
	F_emscripten_builtin_free(m, v135)
	mBase = m.M
	v210 = int32(0)
	v212 = F_FreeFile(m, v35)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L8
	} else {
		goto L47
	}
L47:
	;
	if v212 != 0 {
		v232 = v210
		v235 = v210
		goto L7
	} else {
		goto L48
	}
L48:
	;
	v220 = F_durable_rename(m, int32(247807), int32(119700), int32(15))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L8
	} else {
		goto L49
	}
L49:
	;
	goto L6
L50:
	;
	if v242 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L8
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	F_emscripten_builtin_free(m, v235)
	mBase = m.M
	if v232 != 0 {
		goto L57
	} else {
		goto L58
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(247807)
	F_errmsg(m, int32(314042), v11)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L8
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(518309), int32(824), int32(256723))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L8
	} else {
		goto L56
	}
L56:
	;
	goto L53
L57:
	;
	v263 = F_FreeFile(m, v232)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L8
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v266 = F_unlink(m, int32(247807))
	mBase = m.M
	goto L6
L60:
	;
	goto L59
}
