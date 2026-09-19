package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pgss_ExecutorEnd(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 float64
	_ = v32
	var v35 int32
	_ = v35
	var v36 int64
	_ = v36
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(v5)+8))
	if v6 == int64(0) {
		v54 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_ExecutorEnd[0]))
		if v54 != 0 {
			m.T0[v54].(func(*base.Module, int32))(m, l0)
			mBase = m.M
			v56 = m.ExcPending
			if v56 != 0 {
				return
			} else {
				return
			}
		} else {
			F_standard_ExecutorEnd(m, l0)
			mBase = m.M
			v58 = m.ExcPending
			if v58 != 0 {
				return
			} else {
				return
			}
		}
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
		if v9 == int32(0) {
			v54 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_ExecutorEnd[0]))
			if v54 != 0 {
				m.T0[v54].(func(*base.Module, int32))(m, l0)
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return
				} else {
					return
				}
			} else {
				F_standard_ExecutorEnd(m, l0)
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return
				} else {
					return
				}
			}
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_ExecutorEnd[1]))
			if int32(0) <= v13 {
				v54 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_ExecutorEnd[0]))
				if v54 != 0 {
					m.T0[v54].(func(*base.Module, int32))(m, l0)
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return
					} else {
						return
					}
				} else {
					F_standard_ExecutorEnd(m, l0)
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return
					} else {
						return
					}
				}
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_ExecutorEnd[2]))
				if v17 != int32(2) {
					if v17 != int32(1) {
						v54 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_ExecutorEnd[0]))
						if v54 != 0 {
							m.T0[v54].(func(*base.Module, int32))(m, l0)
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return
							} else {
								return
							}
						} else {
							F_standard_ExecutorEnd(m, l0)
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return
							} else {
								return
							}
						}
					} else {
						v23 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_ExecutorEnd[3]))
						if v23 != 0 {
							v54 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_ExecutorEnd[0]))
							if v54 != 0 {
								m.T0[v54].(func(*base.Module, int32))(m, l0)
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return
								} else {
									return
								}
							} else {
								F_standard_ExecutorEnd(m, l0)
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return
								} else {
									return
								}
							}
						} else {
							F_InstrEndLoop(m, v9)
							mBase = m.M
							v25 = m.ExcPending
							if v25 != 0 {
								return
							} else {
								v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+92))
								v29 = *(*int32)(unsafe.Add(mBase, uint32(v27)+96))
								v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
								v32 = *(*float64)(unsafe.Add(mBase, uint32(v31)+208))
								v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
								v36 = *(*int64)(unsafe.Add(mBase, uint32(v35)+120))
								v41 = *(*int32)(unsafe.Add(mBase, uint32(v35)+180))
								if v41 != 0 {
									v45 = v41 + int32(8)
								} else {
									v45 = int32(0)
								}
								v47 = *(*int32)(unsafe.Add(mBase, uint32(v35)+164))
								v48 = *(*int32)(unsafe.Add(mBase, uint32(v35)+168))
								F_pgss_store(m, v26, v6, v28, v29, int32(1), base.F64_mul(v32, float64(1000)), v36, v31+int32(256), v31+int32(384), v45, int32(0), v47, v48)
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return
								} else {
									v54 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_ExecutorEnd[0]))
									if v54 != 0 {
										m.T0[v54].(func(*base.Module, int32))(m, l0)
										mBase = m.M
										v56 = m.ExcPending
										if v56 != 0 {
											return
										} else {
											return
										}
									} else {
										F_standard_ExecutorEnd(m, l0)
										mBase = m.M
										v58 = m.ExcPending
										if v58 != 0 {
											return
										} else {
											return
										}
									}
								}
							}
						}
					}
				} else {
					F_InstrEndLoop(m, v9)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return
					} else {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+92))
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v27)+96))
						v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
						v32 = *(*float64)(unsafe.Add(mBase, uint32(v31)+208))
						v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						v36 = *(*int64)(unsafe.Add(mBase, uint32(v35)+120))
						v41 = *(*int32)(unsafe.Add(mBase, uint32(v35)+180))
						if v41 != 0 {
							v45 = v41 + int32(8)
						} else {
							v45 = int32(0)
						}
						v47 = *(*int32)(unsafe.Add(mBase, uint32(v35)+164))
						v48 = *(*int32)(unsafe.Add(mBase, uint32(v35)+168))
						F_pgss_store(m, v26, v6, v28, v29, int32(1), base.F64_mul(v32, float64(1000)), v36, v31+int32(256), v31+int32(384), v45, int32(0), v47, v48)
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return
						} else {
							v54 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_ExecutorEnd[0]))
							if v54 != 0 {
								m.T0[v54].(func(*base.Module, int32))(m, l0)
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return
								} else {
									return
								}
							} else {
								F_standard_ExecutorEnd(m, l0)
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return
								} else {
									return
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_pgss_ExecutorRun(m *base.Module, l0 int32, l1 int32, l2 int64) {
	var v7 int32
	_ = v7
	Fn13974(m, l0, l1, l2, int32(_a_F_pgss_ExecutorRun_0), int32(_a_F_pgss_ExecutorRun_1))
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		return
	}
}
func F_pgss_shmem_request(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_shmem_request[0]))
	if v4 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.T0[v4].(func(*base.Module))(m)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_shmem_request[1]))
	v11 = F_hash_estimate_size(m, v9, int32(432))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L4
	} else {
		goto L6
	}
L4:
	;
	return
L5:
	;
	goto L3
L6:
	;
	v13 = F_add_size(m, int32(56), v11)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pgss_shmem_request[2])))
	if v16 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L4
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v32 = int32(_a_F_pgss_shmem_request_0)
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_shmem_request[3]))
	v35 = F_add_size(m, v34, v13)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L4
	} else {
		goto L14
	}
L11:
	;
	F_errmsg_internal(m, int32(_a_F_pgss_shmem_request_1), int32(0))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	F_errfinish(m, int32(_a_F_pgss_shmem_request_2), int32(77), int32(_a_F_pgss_shmem_request_3))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pgss_shmem_request[3])) = v35
	v39 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pgss_shmem_request[2])))
	if v39 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	return
L16:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_shmem_request[4]))
	if v41 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L4
	} else {
		goto L61
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pgss_shmem_request[5])) = int32(16)
	v49 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_shmem_request[6]))
	v51 = F_MemoryContextAlloc(m, v49, int32(1088))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L4
	} else {
		goto L22
	}
L20:
	;
	v54 = v41
	goto L21
L21:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_shmem_request[7]))
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_shmem_request[5]))
	if v58 <= v56 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pgss_shmem_request[4])) = v51
	v54 = v51
	goto L21
L23:
	;
	v60 = int32(1)
	v63 = v56 + v60
	if v63&v56 != 0 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v79 = v54
	v81 = v56
	goto L25
L25:
	;
	v84 = v81*int32(68) + v79
	v85 = int32(_a_F_pgss_shmem_request_4)
	goto L33
L26:
	;
	v68 = v60 << (uint(int32(32)-base.I32_clz(v63)) % 32)
	goto L28
L27:
	;
	v68 = v63
	goto L28
L28:
	;
	v71 = F_repalloc(m, v54, v68*int32(68))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L4
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pgss_shmem_request[5])) = v68
	*(*int32)(unsafe.Add(mBase, _c_F_pgss_shmem_request[4])) = v71
	v78 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_shmem_request[7]))
	v79 = v71
	v81 = v78
	goto L25
L30:
	;
	v205 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v84)+64)) = v205
	v207 = int32(_a_F_pgss_shmem_request_5)
	v209 = *(*int32)(unsafe.Add(mBase, _c_F_pgss_shmem_request[7]))
	*(*int32)(unsafe.Add(mBase, _c_F_pgss_shmem_request[7])) = v209 + v205
	goto L15
L31:
	;
	v202 = F_strlen(m, v191)
	mBase = m.M
	goto L30
L33:
	;
	goto L34
L34:
	;
	v92 = int32(63)
	if (v84^v85)&int32(3) != 0 {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v195 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v192))) = uint8(v195)
	goto L31
L36:
	;
	v176 = v171
	v177 = v172
	v178 = v173
	goto L57
L37:
	;
	if v166 == int32(0) {
		v191 = v164
		v192 = v165
		goto L35
	} else {
		goto L56
	}
L38:
	;
	v164 = v85
	v165 = v84
	v166 = v92
	goto L37
L39:
	;
	goto L40
L40:
	;
	goto L42
L41:
	;
	if base.B2i32(v119 != v120) == int32(0) {
		v191 = v123
		v192 = v117
		goto L35
	} else {
		goto L50
	}
L42:
	;
	v108 = v85
	v109 = v84
	v110 = v92
	goto L45
L45:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108))))
	*(*uint8)(unsafe.Add(mBase, uint32(v109))) = uint8(v112)
	if v112 == int32(0) {
		v171 = v108
		v172 = v109
		v173 = v110
		goto L36
	} else {
		goto L47
	}
L46:
	;
	goto L41
L47:
	;
	v116 = int32(1)
	v117 = v109 + v116
	v119 = v110 - v116
	v120 = int32(0)
	v123 = v108 + v116
	if v123&int32(3) == v120 {
		goto L41
	} else {
		goto L48
	}
L48:
	;
	if v119 != 0 {
		v108 = v123
		v109 = v117
		v110 = v119
		goto L45
	} else {
		goto L49
	}
L49:
	;
	goto L46
L50:
	;
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123))))
	if base.B2i32(v135 == int32(0))|base.B2i32(base.Ui32(v119) < base.Ui32(int32(4))) != 0 {
		v164 = v123
		v165 = v117
		v166 = v119
		goto L37
	} else {
		goto L51
	}
L51:
	;
	v142 = v123
	v143 = v117
	v144 = v119
	goto L52
L52:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v150 = int32(-2139062144)
	if (int32(16843008)-v147|v147)&v150 != v150 {
		v171 = v142
		v172 = v143
		v173 = v144
		goto L36
	} else {
		goto L54
	}
L53:
	;
	v164 = v158
	v165 = v156
	v166 = v160
	goto L37
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v143))) = v147
	v155 = int32(4)
	v156 = v143 + v155
	v158 = v142 + v155
	v160 = v144 - v155
	if base.Ui32(int32(3)) < base.Ui32(v160) {
		v142 = v158
		v143 = v156
		v144 = v160
		goto L52
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	v171 = v164
	v172 = v165
	v173 = v166
	goto L36
L57:
	;
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176))))
	*(*uint8)(unsafe.Add(mBase, uint32(v177))) = uint8(v180)
	if v180 == int32(0) {
		v191 = v176
		v192 = v177
		goto L35
	} else {
		goto L59
	}
L58:
	;
	v191 = v187
	v192 = v185
	goto L35
L59:
	;
	v184 = int32(1)
	v185 = v177 + v184
	v187 = v176 + v184
	v189 = v178 - v184
	if v189 != 0 {
		v176 = v187
		v177 = v185
		v178 = v189
		goto L57
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	F_errmsg_internal(m, int32(_a_F_pgss_shmem_request_6), int32(0))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L4
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_pgss_shmem_request_7), int32(687), int32(_a_F_pgss_shmem_request_8))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L4
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
