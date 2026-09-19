package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_int8_avg_accum(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int64
	_ = v71
	var v72 int32
	_ = v72
	var v78 int64
	_ = v78
	var v81 int64
	_ = v81
	var v82 int64
	_ = v82
	var v87 int64
	_ = v87
	var v88 int64
	_ = v88
	var v91 int64
	_ = v91
	var v94 int64
	_ = v94
	var v95 int64
	_ = v95
	var v102 int64
	_ = v102
	var v113 int64
	_ = v113
	var v114 int64
	_ = v114
	var v115 int64
	_ = v115
	var v119 int64
	_ = v119
	var v120 int64
	_ = v120
	var v126 int64
	_ = v126
	var v127 int64
	_ = v127
	var v129 int64
	_ = v129
	var v135 int64
	_ = v135
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v12 == int32(0) {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v15 != 0 {
			v65 = v15
			v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
			if v67 == int32(0) {
				v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				v71 = *(*int64)(unsafe.Add(mBase, uint32(v70)))
				v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
				if v72 == int32(1) {
					v78 = v71 * (v71 >> (uint(int64(63)) % 64))
					v81 = int64(32)
					v82 = int64(base.Ui64(v71) >> (uint(v81) % 64))
					v87 = int64(4294967295)
					v88 = v71 & v87
					v91 = v88 * v88
					v94 = v88 * v82
					v95 = int64(base.Ui64(v91)>>(uint(v81)%64)) + v94
					v102 = v94 + v95&v87
					*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v78 + v78 + v82*v82 + int64(base.Ui64(v95)>>(uint(v81)%64)) + int64(base.Ui64(v102)>>(uint(v81)%64))
					*(*int64)(unsafe.Add(mBase, uint32(v10))) = v91&v87 | v102<<(uint(v81)%64)
					v113 = *(*int64)(unsafe.Add(mBase, uint32(v65)+32))
					v114 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
					v115 = v113 + v114
					*(*int64)(unsafe.Add(mBase, uint32(v65)+32)) = v115
					v119 = *(*int64)(unsafe.Add(mBase, uint32(v65)+40))
					v120 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v65)+40)) = base.I64_extend_i32_u(base.B2i32(base.Ui64(v115) < base.Ui64(v113))) + (v119 + v120)
				} else {
				}
				v126 = *(*int64)(unsafe.Add(mBase, uint32(v65)+16))
				v127 = v126 + v71
				*(*int64)(unsafe.Add(mBase, uint32(v65)+16)) = v127
				v129 = *(*int64)(unsafe.Add(mBase, uint32(v65)+8))
				*(*int64)(unsafe.Add(mBase, uint32(v65)+8)) = v129 + int64(1)
				v135 = *(*int64)(unsafe.Add(mBase, uint32(v65)+24))
				*(*int64)(unsafe.Add(mBase, uint32(v65)+24)) = base.I64_extend_i32_u(base.B2i32(base.Ui64(v127) < base.Ui64(v126))) + (v135 + v71>>(uint(int64(63))%64))
			} else {
			}
			m.G0 = v10 + int32(32)
			return v65
		} else {
			v18 = v10 + int32(28)
			v19 = int32(0)
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v20 == v19 {
				v37 = int32(0)
				if v18 == v37 {
					v45 = v37
				} else {
					v40 = v37
					v41 = v19
					*(*int32)(unsafe.Add(mBase, uint32(v18))) = v40
					v45 = v41
				}
				v48 = v45
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
				switch v23 - int32(429) {
				case 0:
					if v18 == int32(0) {
						v48 = int32(1)
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v20)+168))
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
						v40 = v30
						v41 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v18))) = v40
						v45 = v41
						v48 = v45
					}
				case 1:
					if v18 == int32(0) {
						v48 = int32(2)
					} else {
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v20)+368))
						v40 = v35
						v41 = int32(2)
						*(*int32)(unsafe.Add(mBase, uint32(v18))) = v40
						v45 = v41
						v48 = v45
					}
				default:
					v37 = int32(0)
					if v18 == v37 {
						v45 = v37
					} else {
						v40 = v37
						v41 = v19
						*(*int32)(unsafe.Add(mBase, uint32(v18))) = v40
						v45 = v41
					}
					v48 = v45
				}
			}
			if v48 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v151 = m.ExcPending
				if v151 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(_a_F_int8_avg_accum_0), int32(0))
					mBase = m.M
					v155 = m.ExcPending
					if v155 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_int8_avg_accum_1), int32(_a_F_int8_avg_accum_2), int32(_a_F_int8_avg_accum_3))
						mBase = m.M
						v160 = m.ExcPending
						if v160 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v51 = int32(_a_F_int8_avg_accum_4)
				v52 = *(*int32)(unsafe.Add(mBase, _c_F_int8_avg_accum[0]))
				v54 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
				*(*int32)(unsafe.Add(mBase, _c_F_int8_avg_accum[0])) = v54
				v57 = F_palloc0(m, int32(48))
				mBase = m.M
				v60 = m.ExcPending
				if v60 != 0 {
					return int32(0)
				} else {
					v61 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v57))) = uint8(v61)
					*(*int32)(unsafe.Add(mBase, _c_F_int8_avg_accum[0])) = v52
					v65 = v57
					v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
					if v67 == int32(0) {
						v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						v71 = *(*int64)(unsafe.Add(mBase, uint32(v70)))
						v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
						if v72 == int32(1) {
							v78 = v71 * (v71 >> (uint(int64(63)) % 64))
							v81 = int64(32)
							v82 = int64(base.Ui64(v71) >> (uint(v81) % 64))
							v87 = int64(4294967295)
							v88 = v71 & v87
							v91 = v88 * v88
							v94 = v88 * v82
							v95 = int64(base.Ui64(v91)>>(uint(v81)%64)) + v94
							v102 = v94 + v95&v87
							*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v78 + v78 + v82*v82 + int64(base.Ui64(v95)>>(uint(v81)%64)) + int64(base.Ui64(v102)>>(uint(v81)%64))
							*(*int64)(unsafe.Add(mBase, uint32(v10))) = v91&v87 | v102<<(uint(v81)%64)
							v113 = *(*int64)(unsafe.Add(mBase, uint32(v65)+32))
							v114 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
							v115 = v113 + v114
							*(*int64)(unsafe.Add(mBase, uint32(v65)+32)) = v115
							v119 = *(*int64)(unsafe.Add(mBase, uint32(v65)+40))
							v120 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v65)+40)) = base.I64_extend_i32_u(base.B2i32(base.Ui64(v115) < base.Ui64(v113))) + (v119 + v120)
						} else {
						}
						v126 = *(*int64)(unsafe.Add(mBase, uint32(v65)+16))
						v127 = v126 + v71
						*(*int64)(unsafe.Add(mBase, uint32(v65)+16)) = v127
						v129 = *(*int64)(unsafe.Add(mBase, uint32(v65)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v65)+8)) = v129 + int64(1)
						v135 = *(*int64)(unsafe.Add(mBase, uint32(v65)+24))
						*(*int64)(unsafe.Add(mBase, uint32(v65)+24)) = base.I64_extend_i32_u(base.B2i32(base.Ui64(v127) < base.Ui64(v126))) + (v135 + v71>>(uint(int64(63))%64))
					} else {
					}
					m.G0 = v10 + int32(32)
					return v65
				}
			}
		}
	} else {
		v18 = v10 + int32(28)
		v19 = int32(0)
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v20 == v19 {
			v37 = int32(0)
			if v18 == v37 {
				v45 = v37
			} else {
				v40 = v37
				v41 = v19
				*(*int32)(unsafe.Add(mBase, uint32(v18))) = v40
				v45 = v41
			}
			v48 = v45
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
			switch v23 - int32(429) {
			case 0:
				if v18 == int32(0) {
					v48 = int32(1)
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v20)+168))
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
					v40 = v30
					v41 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v18))) = v40
					v45 = v41
					v48 = v45
				}
			case 1:
				if v18 == int32(0) {
					v48 = int32(2)
				} else {
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v20)+368))
					v40 = v35
					v41 = int32(2)
					*(*int32)(unsafe.Add(mBase, uint32(v18))) = v40
					v45 = v41
					v48 = v45
				}
			default:
				v37 = int32(0)
				if v18 == v37 {
					v45 = v37
				} else {
					v40 = v37
					v41 = v19
					*(*int32)(unsafe.Add(mBase, uint32(v18))) = v40
					v45 = v41
				}
				v48 = v45
			}
		}
		if v48 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v151 = m.ExcPending
			if v151 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_int8_avg_accum_0), int32(0))
				mBase = m.M
				v155 = m.ExcPending
				if v155 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_int8_avg_accum_1), int32(_a_F_int8_avg_accum_2), int32(_a_F_int8_avg_accum_3))
					mBase = m.M
					v160 = m.ExcPending
					if v160 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v51 = int32(_a_F_int8_avg_accum_4)
			v52 = *(*int32)(unsafe.Add(mBase, _c_F_int8_avg_accum[0]))
			v54 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
			*(*int32)(unsafe.Add(mBase, _c_F_int8_avg_accum[0])) = v54
			v57 = F_palloc0(m, int32(48))
			mBase = m.M
			v60 = m.ExcPending
			if v60 != 0 {
				return int32(0)
			} else {
				v61 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v57))) = uint8(v61)
				*(*int32)(unsafe.Add(mBase, _c_F_int8_avg_accum[0])) = v52
				v65 = v57
				v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
				if v67 == int32(0) {
					v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					v71 = *(*int64)(unsafe.Add(mBase, uint32(v70)))
					v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
					if v72 == int32(1) {
						v78 = v71 * (v71 >> (uint(int64(63)) % 64))
						v81 = int64(32)
						v82 = int64(base.Ui64(v71) >> (uint(v81) % 64))
						v87 = int64(4294967295)
						v88 = v71 & v87
						v91 = v88 * v88
						v94 = v88 * v82
						v95 = int64(base.Ui64(v91)>>(uint(v81)%64)) + v94
						v102 = v94 + v95&v87
						*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v78 + v78 + v82*v82 + int64(base.Ui64(v95)>>(uint(v81)%64)) + int64(base.Ui64(v102)>>(uint(v81)%64))
						*(*int64)(unsafe.Add(mBase, uint32(v10))) = v91&v87 | v102<<(uint(v81)%64)
						v113 = *(*int64)(unsafe.Add(mBase, uint32(v65)+32))
						v114 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
						v115 = v113 + v114
						*(*int64)(unsafe.Add(mBase, uint32(v65)+32)) = v115
						v119 = *(*int64)(unsafe.Add(mBase, uint32(v65)+40))
						v120 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v65)+40)) = base.I64_extend_i32_u(base.B2i32(base.Ui64(v115) < base.Ui64(v113))) + (v119 + v120)
					} else {
					}
					v126 = *(*int64)(unsafe.Add(mBase, uint32(v65)+16))
					v127 = v126 + v71
					*(*int64)(unsafe.Add(mBase, uint32(v65)+16)) = v127
					v129 = *(*int64)(unsafe.Add(mBase, uint32(v65)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v65)+8)) = v129 + int64(1)
					v135 = *(*int64)(unsafe.Add(mBase, uint32(v65)+24))
					*(*int64)(unsafe.Add(mBase, uint32(v65)+24)) = base.I64_extend_i32_u(base.B2i32(base.Ui64(v127) < base.Ui64(v126))) + (v135 + v71>>(uint(int64(63))%64))
				} else {
				}
				m.G0 = v10 + int32(32)
				return v65
			}
		}
	}
}
func F_int8_avg_combine(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int64
	_ = v105
	var v107 int64
	_ = v107
	var v108 int64
	_ = v108
	var v113 int64
	_ = v113
	var v116 int64
	_ = v116
	var v119 int64
	_ = v119
	var v120 int64
	_ = v120
	var v121 int64
	_ = v121
	var v122 int64
	_ = v122
	var v126 int64
	_ = v126
	var v130 int32
	_ = v130
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = v11 + int32(8)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v16 == v2 {
		v33 = int32(0)
		if v14 == v33 {
			v41 = v33
		} else {
			v36 = v33
			v37 = v2
			*(*int32)(unsafe.Add(mBase, uint32(v14))) = v36
			v41 = v37
		}
		v44 = v41
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
		switch v19 - int32(429) {
		case 0:
			if v14 == int32(0) {
				v44 = int32(1)
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v16)+168))
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
				v36 = v26
				v37 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v14))) = v36
				v41 = v37
				v44 = v41
			}
		case 1:
			if v14 == int32(0) {
				v44 = int32(2)
			} else {
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v16)+368))
				v36 = v31
				v37 = int32(2)
				*(*int32)(unsafe.Add(mBase, uint32(v14))) = v36
				v41 = v37
				v44 = v41
			}
		default:
			v33 = int32(0)
			if v14 == v33 {
				v41 = v33
			} else {
				v36 = v33
				v37 = v2
				*(*int32)(unsafe.Add(mBase, uint32(v14))) = v36
				v41 = v37
			}
			v44 = v41
		}
	}
	if v44 != 0 {
		v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
		if v45 == int32(0) {
			v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v49 = v48
		} else {
			v49 = v2
		}
		v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
		if v50 != 0 {
			v130 = v49
			m.G0 = v11 + int32(16)
			return v130
		} else {
			v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			if v51 == int32(0) {
				v130 = v49
				m.G0 = v11 + int32(16)
				return v130
			} else {
				if v49 == int32(0) {
					v56 = int32(_a_F_int8_avg_combine_0)
					v57 = *(*int32)(unsafe.Add(mBase, _c_F_int8_avg_combine[0]))
					v59 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
					*(*int32)(unsafe.Add(mBase, _c_F_int8_avg_combine[0])) = v59
					v62 = v11 + int32(12)
					v63 = int32(0)
					v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					if v64 == v63 {
						v81 = int32(0)
						if v62 == v81 {
							v89 = v81
						} else {
							v84 = v81
							v85 = v63
							*(*int32)(unsafe.Add(mBase, uint32(v62))) = v84
							v89 = v85
						}
						v92 = v89
					} else {
						v67 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
						switch v67 - int32(429) {
						case 0:
							if v62 == int32(0) {
								v92 = int32(1)
							} else {
								v73 = *(*int32)(unsafe.Add(mBase, uint32(v64)+168))
								v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+20))
								v84 = v74
								v85 = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v62))) = v84
								v89 = v85
								v92 = v89
							}
						case 1:
							if v62 == int32(0) {
								v92 = int32(2)
							} else {
								v79 = *(*int32)(unsafe.Add(mBase, uint32(v64)+368))
								v84 = v79
								v85 = int32(2)
								*(*int32)(unsafe.Add(mBase, uint32(v62))) = v84
								v89 = v85
								v92 = v89
							}
						default:
							v81 = int32(0)
							if v62 == v81 {
								v89 = v81
							} else {
								v84 = v81
								v85 = v63
								*(*int32)(unsafe.Add(mBase, uint32(v62))) = v84
								v89 = v85
							}
							v92 = v89
						}
					}
					if v92 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v156 = m.ExcPending
						if v156 != 0 {
							return int32(0)
						} else {
							F_errmsg_internal(m, int32(_a_F_int8_avg_combine_1), int32(0))
							mBase = m.M
							v160 = m.ExcPending
							if v160 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_int8_avg_combine_2), int32(_a_F_int8_avg_combine_3), int32(_a_F_int8_avg_combine_4))
								mBase = m.M
								v165 = m.ExcPending
								if v165 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v96 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
						*(*int32)(unsafe.Add(mBase, _c_F_int8_avg_combine[0])) = v96
						v99 = F_palloc0(m, int32(48))
						mBase = m.M
						v102 = m.ExcPending
						if v102 != 0 {
							return int32(0)
						} else {
							v103 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v99))) = uint8(v103)
							v105 = *(*int64)(unsafe.Add(mBase, uint32(v51)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v99)+8)) = v105
							v107 = *(*int64)(unsafe.Add(mBase, uint32(v51)+16))
							v108 = *(*int64)(unsafe.Add(mBase, uint32(v51)+24))
							*(*int64)(unsafe.Add(mBase, uint32(v99)+24)) = v108
							*(*int64)(unsafe.Add(mBase, uint32(v99)+16)) = v107
							*(*int32)(unsafe.Add(mBase, _c_F_int8_avg_combine[0])) = v57
							v130 = v99
							m.G0 = v11 + int32(16)
							return v130
						}
					}
				} else {
					v113 = *(*int64)(unsafe.Add(mBase, uint32(v51)+8))
					if v113 <= int64(0) {
						v130 = v49
					} else {
						v116 = *(*int64)(unsafe.Add(mBase, uint32(v49)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v49)+8)) = v116 + v113
						v119 = *(*int64)(unsafe.Add(mBase, uint32(v51)+24))
						v120 = *(*int64)(unsafe.Add(mBase, uint32(v49)+16))
						v121 = *(*int64)(unsafe.Add(mBase, uint32(v51)+16))
						v122 = v120 + v121
						*(*int64)(unsafe.Add(mBase, uint32(v49)+16)) = v122
						v126 = *(*int64)(unsafe.Add(mBase, uint32(v49)+24))
						*(*int64)(unsafe.Add(mBase, uint32(v49)+24)) = base.I64_extend_i32_u(base.B2i32(base.Ui64(v122) < base.Ui64(v120))) + (v119 + v126)
						v130 = v49
					}
					m.G0 = v11 + int32(16)
					return v130
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v143 = m.ExcPending
		if v143 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_int8_avg_combine_1), int32(0))
			mBase = m.M
			v147 = m.ExcPending
			if v147 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_int8_avg_combine_2), int32(_a_F_int8_avg_combine_5), int32(_a_F_int8_avg_combine_6))
				mBase = m.M
				v152 = m.ExcPending
				if v152 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_int8_dist(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn13880(m, l0, int32(_a_F_int8_dist_0), int32(_a_F_int8_dist_1), int32(_a_F_int8_dist_2))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_int8_to_char(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
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
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v180 int32
	_ = v180
	var v182 float64
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int64
	_ = v195
	var v196 int32
	_ = v196
	var v200 int64
	_ = v200
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v360 int32
	_ = v360
	v2 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(48)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v19 = F_pg_detoast_datum_packed(m, v18)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v14 + int32(48)
	return v360
L2:
	;
	if base.Ui32(int32(268435454)) <= base.Ui32(v52-int32(1)) {
		goto L15
	} else {
		goto L16
	}
L3:
	;
	return int32(0)
L4:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if v23 == int32(1) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	if v29 == int32(18) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v40 = int32(1)
	if v23&v40 != 0 {
		v52 = int32(base.Ui32(v23)>>(uint(v40)%32)) - v40
		goto L2
	} else {
		goto L14
	}
L8:
	;
	v32 = int32(16)
	goto L10
L9:
	;
	v32 = int32(0)
	goto L10
L10:
	;
	if base.Ui32((v29-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v39 = int32(4)
	goto L13
L12:
	;
	v39 = v32
	goto L13
L13:
	;
	v52 = v39
	goto L2
L14:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v52 = int32(base.Ui32(v46)>>(uint(int32(2))%32)) - int32(4)
	goto L2
L15:
	;
	v58 = F_cstring_to_text(m, int32(_a_F_int8_to_char_0))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L3
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v64 = F_palloc0(m, v52<<(uint(int32(3))%32)|int32(5))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L3
	} else {
		goto L19
	}
L18:
	;
	v360 = v58
	goto L1
L19:
	;
	v70 = F_NUM_cache(m, v52, v14+int32(12), v19, v14+int32(11))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L3
	} else {
		goto L20
	}
L20:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
	if v72&int32(1024) != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v339 = v64 + int32(4)
	F_NUM_processor(m, v70, v14+int32(12), v339, v328, int32(0), v331, v333, int32(1))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L3
	} else {
		goto L107
	}
L22:
	;
	if base.Ui64(int64(4294967296)) <= base.Ui64(v17+int64(2147483648)) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	if v72&int32(_a_F_int8_to_char_1) != 0 {
		goto L29
	} else {
		goto L30
	}
L25:
	;
	v81 = int32(2147483647)
	goto L27
L26:
	;
	v81 = base.I32_wrap_i64(v17)
	goto L27
L27:
	;
	v82 = F_int_to_roman(m, v81)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L3
	} else {
		goto L28
	}
L28:
	;
	v328 = v82
	v331 = v2
	v333 = v2
	goto L21
L29:
	;
	v86 = F_int64_to_numeric(m, v17)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L3
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	if v72&int32(2048) != 0 {
		goto L59
	} else {
		goto L60
	}
L32:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	v89 = F_numeric_out_sci(m, v86, v88)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L3
	} else {
		goto L33
	}
L33:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
	if v91 == int32(45) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v328 = v89
	v331 = v2
	v333 = v2
	goto L21
L35:
	;
	goto L36
L36:
	;
	v94 = F_strlen(m, v89)
	mBase = m.M
	v97 = F_palloc(m, v94+int32(2))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L3
	} else {
		goto L37
	}
L37:
	;
	v99 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v97))) = uint8(v99)
	v102 = v97 + int32(1)
	if (v89^v102)&int32(3) != 0 {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	v328 = v97
	v331 = v2
	v333 = v2
	goto L21
L39:
	;
	goto L38
L40:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v157))) = uint8(v156)
	if v156&int32(255) == int32(0) {
		goto L39
	} else {
		goto L55
	}
L41:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
	v155 = v89
	v156 = v108
	v157 = v102
	goto L40
L42:
	;
	goto L43
L43:
	;
	if v89&int32(3) != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v112 = v89
	v114 = v102
	goto L47
L45:
	;
	v126 = v89
	v128 = v102
	goto L46
L46:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	v133 = int32(-2139062144)
	if (int32(16843008)-v130|v130)&v133 != v133 {
		v155 = v126
		v156 = v130
		v157 = v128
		goto L40
	} else {
		goto L51
	}
L47:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112))))
	*(*uint8)(unsafe.Add(mBase, uint32(v114))) = uint8(v115)
	if v115 == int32(0) {
		goto L39
	} else {
		goto L49
	}
L48:
	;
	v126 = v122
	v128 = v120
	goto L46
L49:
	;
	v119 = int32(1)
	v120 = v114 + v119
	v122 = v112 + v119
	if v122&int32(3) != 0 {
		v112 = v122
		v114 = v120
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	v138 = v126
	v139 = v130
	v140 = v128
	goto L52
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v140))) = v139
	v142 = int32(4)
	v143 = v140 + v142
	v145 = v138 + v142
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
	v150 = int32(-2139062144)
	if (int32(16843008)-v147|v147)&v150 == v150 {
		v138 = v145
		v139 = v147
		v140 = v143
		goto L52
	} else {
		goto L54
	}
L53:
	;
	v155 = v145
	v156 = v147
	v157 = v143
	goto L40
L54:
	;
	goto L53
L55:
	;
	v164 = v155
	v166 = v157
	goto L56
L56:
	;
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v166)+1)) = uint8(v167)
	v169 = int32(1)
	if v167 != 0 {
		v164 = v164 + v169
		v166 = v166 + v169
		goto L56
	} else {
		goto L58
	}
L57:
	;
	goto L39
L58:
	;
	goto L57
L59:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
	v182 = F_pow(m, float64(10), base.F64_convert_i32_s(v180))
	mBase = m.M
	v185 = F_Int64GetDatum(m, v17)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L3
	} else {
		goto L62
	}
L60:
	;
	v200 = v17
	goto L61
L61:
	;
	v204 = F_Int64GetDatum(m, v200)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L3
	} else {
		goto L66
	}
L62:
	;
	v189 = F_Float8GetDatum(m, v182)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L3
	} else {
		goto L63
	}
L63:
	;
	v191 = F_DirectFunctionCall1Coll(m, int32(1299), int32(0), v189)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L3
	} else {
		goto L64
	}
L64:
	;
	v193 = F_DirectFunctionCall2Coll(m, int32(1263), int32(0), v185, v191)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L3
	} else {
		goto L65
	}
L65:
	;
	v195 = *(*int64)(unsafe.Add(mBase, uint32(v193)))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v180 + v196
	v200 = v195
	goto L61
L66:
	;
	v206 = F_DirectFunctionCall1Coll(m, int32(1300), int32(0), v204)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L3
	} else {
		goto L67
	}
L67:
	;
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206))))
	v210 = base.B2i32(v208 == int32(45))
	v211 = v206 + v210
	v212 = F_strlen(m, v211)
	mBase = m.M
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	if v213 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v217 = F_palloc(m, v212+v213+int32(2))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L3
	} else {
		goto L71
	}
L69:
	;
	v303 = v211
	goto L70
L70:
	;
	if v208 == int32(45) {
		goto L96
	} else {
		goto L97
	}
L71:
	;
	if (v211^v217)&int32(3) != 0 {
		goto L75
	} else {
		goto L76
	}
L72:
	;
	v293 = v217 + v212
	v294 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v293))) = uint8(v294)
	if v213 != 0 {
		goto L93
	} else {
		goto L94
	}
L73:
	;
	goto L72
L74:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v273))) = uint8(v272)
	if v272&int32(255) == int32(0) {
		goto L73
	} else {
		goto L89
	}
L75:
	;
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211))))
	v271 = v211
	v272 = v224
	v273 = v217
	goto L74
L76:
	;
	goto L77
L77:
	;
	if v211&int32(3) != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v228 = v211
	v230 = v217
	goto L81
L79:
	;
	v242 = v211
	v244 = v217
	goto L80
L80:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v242)))
	v249 = int32(-2139062144)
	if (int32(16843008)-v246|v246)&v249 != v249 {
		v271 = v242
		v272 = v246
		v273 = v244
		goto L74
	} else {
		goto L85
	}
L81:
	;
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228))))
	*(*uint8)(unsafe.Add(mBase, uint32(v230))) = uint8(v231)
	if v231 == int32(0) {
		goto L73
	} else {
		goto L83
	}
L82:
	;
	v242 = v238
	v244 = v236
	goto L80
L83:
	;
	v235 = int32(1)
	v236 = v230 + v235
	v238 = v228 + v235
	if v238&int32(3) != 0 {
		v228 = v238
		v230 = v236
		goto L81
	} else {
		goto L84
	}
L84:
	;
	goto L82
L85:
	;
	v254 = v242
	v255 = v246
	v256 = v244
	goto L86
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v256))) = v255
	v258 = int32(4)
	v259 = v256 + v258
	v261 = v254 + v258
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v254)+4))
	v266 = int32(-2139062144)
	if (int32(16843008)-v263|v263)&v266 == v266 {
		v254 = v261
		v255 = v263
		v256 = v259
		goto L86
	} else {
		goto L88
	}
L87:
	;
	v271 = v261
	v272 = v263
	v273 = v259
	goto L74
L88:
	;
	goto L87
L89:
	;
	v280 = v271
	v282 = v273
	goto L90
L90:
	;
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v280)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v282)+1)) = uint8(v283)
	v285 = int32(1)
	if v283 != 0 {
		v280 = v280 + v285
		v282 = v282 + v285
		goto L90
	} else {
		goto L92
	}
L91:
	;
	goto L73
L92:
	;
	goto L91
L93:
	;
	base.MemoryFill(m, v293+int32(1), int32(48), v213)
	goto L95
L94:
	;
	goto L95
L95:
	;
	v301 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v293+v213)+1)) = uint8(v301)
	v303 = v217
	goto L70
L96:
	;
	v307 = int32(45)
	goto L98
L97:
	;
	v307 = int32(43)
	goto L98
L98:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	if v212 < v308 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v328 = v303
	v331 = v308 - v212
	v333 = v307
	goto L21
L100:
	;
	goto L101
L101:
	;
	v311 = int32(0)
	if v212 <= v308 {
		v328 = v303
		v331 = v311
		v333 = v307
		goto L21
	} else {
		goto L102
	}
L102:
	;
	v313 = v308 + v213
	v316 = F_palloc(m, v313+int32(2))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L3
	} else {
		goto L103
	}
L103:
	;
	v319 = v313 + int32(1)
	if v319 != 0 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	base.MemoryFill(m, v316, int32(35), v319)
	goto L106
L105:
	;
	goto L106
L106:
	;
	v323 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v316+v319))) = uint8(v323)
	v326 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v316+v308))) = uint8(v326)
	v328 = v316
	v331 = v311
	v333 = v307
	goto L21
L107:
	;
	v344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+11)))
	if v344 == int32(1) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	F_pfree(m, v70)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L3
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v349 = F_strlen(m, v339)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v349<<(uint(int32(2))%32) + int32(16)
	v360 = v64
	goto L1
L111:
	;
	goto L110
}
