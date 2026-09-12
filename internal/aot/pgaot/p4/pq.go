package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pq_getkeepalivesinterval(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	v2 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	if l0 == v2 {
		v22 = v2
	} else {
		v10 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
		if v10 == int32(1) {
			v22 = v2
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+404))
			if v13 != 0 {
				v22 = v13
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+388))
				if v14 != 0 {
					v22 = v14
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = int32(4)
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(388))))
					v22 = v20
				}
			}
		}
	}
	m.G0 = v6 + int32(16)
	return v22
}
func F_pq_getmsgbytes(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	if int32(0) <= l1 {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if l1 <= v6-v7 {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l1 + v7
			v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			return v31 + v7
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(16908800))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(404472), int32(0))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(493474), int32(515), int32(158687))
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
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
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(16908800))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(404472), int32(0))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(493474), int32(515), int32(158687))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
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
}
func F_pq_getmsgstring(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = v5 + v6
	v8 = F_strlen(m, v7)
	mBase = m.M
	v9 = v8 + v5
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v10 <= v9 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(16908800))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(404530), int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(493474), int32(595), int32(329656))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v9 + int32(1)
		v33 = F_pg_client_to_server(m, v7, v8)
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return int32(0)
		} else {
			return v33
		}
	}
}
func F_pq_getmsgtext(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	if l1 < int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v38 = m.ExcPending
		if v38 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(16908800))
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(404472), int32(0))
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(493474), int32(554), int32(62616))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v7-v8 < l1 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(16908800))
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(404472), int32(0))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(493474), int32(554), int32(62616))
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l1 + v8
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v14 = v13 + v8
			v15 = F_pg_client_to_server(m, v14, l1)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				if v15 != v14 {
					v20 = F_strlen(m, v15)
					mBase = m.M
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v20
					return v15
				} else {
					v25 = F_palloc(m, l1+int32(1))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						if l1 != 0 {
							v27 = F__emscripten_memcpy_bulkmem(m, v25, v14, l1)
							mBase = m.M
							v28 = v27
						} else {
							v28 = v25
						}
						v30 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v28+l1))) = uint8(v30)
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = l1
						return v28
					}
				}
			}
		}
	}
}
func F_pq_init(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	v4 = m.G0
	v6 = v4 + int32(-64)
	m.G0 = v6
	v9 = F_palloc0(m, int32(524))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		*(*int32)(unsafe.Add(mBase, uint32(v9))) = v13
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
		if v19 != 0 {
			v20 = F__emscripten_memcpy_bulkmem(m, v9+int32(144), l0+int32(4), v19)
			mBase = m.M
		} else {
		}
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
		*(*int32)(unsafe.Add(mBase, uint32(v9)+140)) = int32(128)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+272)) = v22
		v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9+int32(12)))))
		if v28 != int32(1) {
			v31 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v6)+60)) = v31
			*(*int32)(unsafe.Add(mBase, uint32(v6)+60)) = v31
			v38 = *(*int32)(unsafe.Add(mBase, _consts[585]))
			v40 = m.G0
			v42 = v40 - int32(16)
			m.G0 = v42
			*(*int32)(unsafe.Add(mBase, uint32(v42)+12)) = v38
			if v9 == int32(0) {
			} else {
				v47 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+12)))
				if v47 == int32(1) {
				} else {
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v9)+400))
					if v38 == v50 {
					} else {
						v52 = *(*int32)(unsafe.Add(mBase, uint32(v9)+384))
						if int32(0) < v52 {
							if v38 == int32(0) {
								v60 = *(*int32)(unsafe.Add(mBase, uint32(v9)+384))
								*(*int32)(unsafe.Add(mBase, uint32(v42)+12)) = v60
							} else {
							}
							v63 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v9)+400)) = v63
						} else {
							v55 = F_pq_getkeepalivesidle(m, v9)
							mBase = m.M
							if int32(0) <= v55 {
								if v38 == int32(0) {
									v60 = *(*int32)(unsafe.Add(mBase, uint32(v9)+384))
									*(*int32)(unsafe.Add(mBase, uint32(v42)+12)) = v60
								} else {
								}
								v63 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
								*(*int32)(unsafe.Add(mBase, uint32(v9)+400)) = v63
							} else {
							}
						}
					}
				}
			}
			m.G0 = v42 + int32(16)
			v69 = *(*int32)(unsafe.Add(mBase, _consts[586]))
			v71 = m.G0
			v73 = v71 - int32(16)
			m.G0 = v73
			*(*int32)(unsafe.Add(mBase, uint32(v73)+12)) = v69
			if v9 == int32(0) {
			} else {
				v78 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+12)))
				if v78 == int32(1) {
				} else {
					v81 = *(*int32)(unsafe.Add(mBase, uint32(v9)+404))
					if v69 == v81 {
					} else {
						v83 = *(*int32)(unsafe.Add(mBase, uint32(v9)+388))
						if int32(0) < v83 {
							if v69 == int32(0) {
								v91 = *(*int32)(unsafe.Add(mBase, uint32(v9)+388))
								*(*int32)(unsafe.Add(mBase, uint32(v73)+12)) = v91
							} else {
							}
							v94 = *(*int32)(unsafe.Add(mBase, uint32(v73)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v9)+404)) = v94
						} else {
							v86 = F_pq_getkeepalivesinterval(m, v9)
							mBase = m.M
							if int32(0) <= v86 {
								if v69 == int32(0) {
									v91 = *(*int32)(unsafe.Add(mBase, uint32(v9)+388))
									*(*int32)(unsafe.Add(mBase, uint32(v73)+12)) = v91
								} else {
								}
								v94 = *(*int32)(unsafe.Add(mBase, uint32(v73)+12))
								*(*int32)(unsafe.Add(mBase, uint32(v9)+404)) = v94
							} else {
							}
						}
					}
				}
			}
			m.G0 = v73 + int32(16)
			v100 = *(*int32)(unsafe.Add(mBase, _consts[587]))
			v102 = m.G0
			v104 = v102 - int32(16)
			m.G0 = v104
			*(*int32)(unsafe.Add(mBase, uint32(v104)+12)) = v100
			if v9 == int32(0) {
			} else {
				v109 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+12)))
				if v109 == int32(1) {
				} else {
					v112 = *(*int32)(unsafe.Add(mBase, uint32(v9)+408))
					if v100 == v112 {
					} else {
						v114 = *(*int32)(unsafe.Add(mBase, uint32(v9)+392))
						if int32(0) < v114 {
							if v100 == int32(0) {
								v122 = *(*int32)(unsafe.Add(mBase, uint32(v9)+392))
								*(*int32)(unsafe.Add(mBase, uint32(v104)+12)) = v122
							} else {
							}
							v125 = *(*int32)(unsafe.Add(mBase, uint32(v104)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v9)+408)) = v125
						} else {
							v117 = F_pq_getkeepalivescount(m, v9)
							mBase = m.M
							if int32(0) <= v117 {
								if v100 == int32(0) {
									v122 = *(*int32)(unsafe.Add(mBase, uint32(v9)+392))
									*(*int32)(unsafe.Add(mBase, uint32(v104)+12)) = v122
								} else {
								}
								v125 = *(*int32)(unsafe.Add(mBase, uint32(v104)+12))
								*(*int32)(unsafe.Add(mBase, uint32(v9)+408)) = v125
							} else {
							}
						}
					}
				}
			}
			m.G0 = v104 + int32(16)
			v131 = *(*int32)(unsafe.Add(mBase, _consts[588]))
			v133 = m.G0
			v135 = v133 - int32(16)
			m.G0 = v135
			*(*int32)(unsafe.Add(mBase, uint32(v135)+12)) = v131
			if v9 == int32(0) {
			} else {
				v140 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+12)))
				if v140 == int32(1) {
				} else {
					v143 = *(*int32)(unsafe.Add(mBase, uint32(v9)+412))
					if v131 == v143 {
					} else {
						v145 = *(*int32)(unsafe.Add(mBase, uint32(v9)+396))
						if int32(0) < v145 {
							if v131 == int32(0) {
								v153 = *(*int32)(unsafe.Add(mBase, uint32(v9)+396))
								*(*int32)(unsafe.Add(mBase, uint32(v135)+12)) = v153
							} else {
							}
							v156 = *(*int32)(unsafe.Add(mBase, uint32(v135)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v9)+412)) = v156
						} else {
							v148 = F_pq_gettcpusertimeout(m, v9)
							mBase = m.M
							if int32(0) <= v148 {
								if v131 == int32(0) {
									v153 = *(*int32)(unsafe.Add(mBase, uint32(v9)+396))
									*(*int32)(unsafe.Add(mBase, uint32(v135)+12)) = v153
								} else {
								}
								v156 = *(*int32)(unsafe.Add(mBase, uint32(v135)+12))
								*(*int32)(unsafe.Add(mBase, uint32(v9)+412)) = v156
							} else {
							}
						}
					}
				}
			}
			m.G0 = v135 + int32(16)
		} else {
		}
		v162 = int32(8192)
		*(*int32)(unsafe.Add(mBase, _consts[589])) = v162
		v166 = *(*int32)(unsafe.Add(mBase, _consts[36]))
		v168 = F_MemoryContextAlloc(m, v166, v162)
		mBase = m.M
		v169 = m.ExcPending
		if v169 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[590])) = v168
			v172 = int32(0)
			*(*int32)(unsafe.Add(mBase, _consts[591])) = v172
			*(*int32)(unsafe.Add(mBase, _consts[592])) = v172
			*(*int32)(unsafe.Add(mBase, _consts[593])) = v172
			*(*int32)(unsafe.Add(mBase, _consts[594])) = v172
			*(*uint8)(unsafe.Add(mBase, _consts[595])) = uint8(v172)
			*(*uint8)(unsafe.Add(mBase, _consts[596])) = uint8(v172)
			F_on_proc_exit(m, int32(801))
			mBase = m.M
			v191 = m.ExcPending
			if v191 != 0 {
				return int32(0)
			} else {
				v194 = m.G0
				v196 = v194 - int32(16)
				m.G0 = v196
				*(*int32)(unsafe.Add(mBase, uint32(v196))) = int32(2048)
				m.G0 = v196 + int32(16)
				*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = int32(1)
				v217 = F_CreateWaitEventSet(m, int32(0), int32(3))
				mBase = m.M
				v218 = m.ExcPending
				if v218 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, _consts[597])) = v217
					v221 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
					F_AddWaitEventToSet(m, v217, int32(4), v221, int32(0))
					mBase = m.M
					v224 = m.ExcPending
					if v224 != 0 {
						return int32(0)
					} else {
						v226 = *(*int32)(unsafe.Add(mBase, _consts[597]))
						v230 = *(*int32)(unsafe.Add(mBase, _consts[506]))
						F_AddWaitEventToSet(m, v226, int32(1), int32(-1), v230)
						mBase = m.M
						v232 = m.ExcPending
						if v232 != 0 {
							return int32(0)
						} else {
							v234 = *(*int32)(unsafe.Add(mBase, _consts[597]))
							F_AddWaitEventToSet(m, v234, int32(16), int32(-1), int32(0))
							mBase = m.M
							v239 = m.ExcPending
							if v239 != 0 {
								return int32(0)
							} else {
								m.G0 = v6 - int32(-64)
								return v9
							}
						}
					}
				}
			}
		}
	}
}
func F_pq_parse_errornotice(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v302 int32
	_ = v302
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v330 int32
	_ = v330
	var v338 int32
	_ = v338
	var v346 int32
	_ = v346
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v423 int32
	_ = v423
	var v430 int32
	_ = v430
	var v436 int32
	_ = v436
	var v441 int32
	_ = v441
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	if l1&int32(3) == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(21)
	v42 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+96)) = v42
	v44 = F_pq_getmsgbyte(m, l0)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L12
	} else {
		goto L13
	}
L2:
	;
	v15 = l1 + int32(100)
	if base.Ui32(v15) <= base.Ui32(l1) {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v36 = F__emscripten_memset_bulkmem(m, l1+int32(4), base.I32_extend8_s(int32(0)), int32(96))
	mBase = m.M
	goto L10
L5:
	;
	v21 = l1 + int32(4)
	if base.Ui32(v21) < base.Ui32(v15) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v23 = v15
	goto L8
L7:
	;
	v23 = v21
	goto L8
L8:
	;
	v30 = F__emscripten_memset_bulkmem(m, l1, base.I32_extend8_s(int32(0)), (l1^int32(-1)+v23)&int32(-4)+int32(4))
	mBase = m.M
	goto L9
L9:
	;
	goto L1
L10:
	;
	goto L1
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L12
	} else {
		goto L152
	}
L12:
	;
	return
L13:
	;
	v47 = v44 << (uint(int32(24)) % 32)
	if v47 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v52 = v47
	goto L17
L15:
	;
	goto L16
L16:
	;
	F_pq_getmsgend(m, l0)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L12
	} else {
		goto L151
	}
L17:
	;
	v53 = F_pq_getmsgrawstring(m, l0)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L12
	} else {
		goto L19
	}
L18:
	;
	goto L16
L19:
	;
	v56 = v52 >> (uint(int32(24)) % 32)
	switch v56 - int32(67) {
	case 0:
		goto L37
	case 1:
		goto L36
	default:
		goto L22
	case 3:
		goto L25
	case 5:
		goto L35
	case 9:
		goto L24
	case 10:
		goto L21
	case 13:
		goto L34
	case 15:
		goto L23
	case 16:
		goto L20
	case 19:
		goto L38
	case 20:
		goto L31
	case 32:
		goto L28
	case 33:
		goto L27
	case 43:
		goto L26
	case 45:
		goto L33
	case 46:
		goto L32
	case 48:
		goto L30
	case 49:
		goto L29
	}
L20:
	;
	v413 = F_pq_getmsgbyte(m, l0)
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L12
	} else {
		goto L149
	}
L21:
	;
	v410 = F_pstrdup(m, v53)
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L12
	} else {
		goto L148
	}
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L12
	} else {
		goto L145
	}
L23:
	;
	v394 = F_pstrdup(m, v53)
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L12
	} else {
		goto L144
	}
L24:
	;
	v391 = F_pg_strtoint32(m, v53)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L12
	} else {
		goto L143
	}
L25:
	;
	v388 = F_pstrdup(m, v53)
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L12
	} else {
		goto L142
	}
L26:
	;
	v385 = F_pstrdup(m, v53)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L12
	} else {
		goto L141
	}
L27:
	;
	v382 = F_pstrdup(m, v53)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L12
	} else {
		goto L140
	}
L28:
	;
	v379 = F_pstrdup(m, v53)
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L12
	} else {
		goto L139
	}
L29:
	;
	v376 = F_pstrdup(m, v53)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L12
	} else {
		goto L138
	}
L30:
	;
	v373 = F_pstrdup(m, v53)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L12
	} else {
		goto L137
	}
L31:
	;
	v370 = F_pstrdup(m, v53)
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L12
	} else {
		goto L136
	}
L32:
	;
	v367 = F_pstrdup(m, v53)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L12
	} else {
		goto L135
	}
L33:
	;
	v364 = F_pg_strtoint32(m, v53)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L12
	} else {
		goto L134
	}
L34:
	;
	v361 = F_pg_strtoint32(m, v53)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L12
	} else {
		goto L133
	}
L35:
	;
	v358 = F_pstrdup(m, v53)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L12
	} else {
		goto L132
	}
L36:
	;
	v355 = F_pstrdup(m, v53)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L12
	} else {
		goto L131
	}
L37:
	;
	v314 = F_strlen(m, v53)
	mBase = m.M
	if v314 != int32(5) {
		goto L11
	} else {
		goto L130
	}
L38:
	;
	v59 = int32(536008)
	v62 = int32(*(*uint8)(unsafe.Add(mBase, _consts[601])))
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v63 == int32(0) {
		v82 = v62
		v83 = v63
		goto L40
	} else {
		goto L41
	}
L39:
	;
	if v83-v82 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L40:
	;
	goto L39
L41:
	;
	if v62 != v63 {
		v82 = v62
		v83 = v63
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v67 = v53
	v68 = v59
	goto L43
L43:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+1)))
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+1)))
	if v72 == int32(0) {
		v82 = v71
		v83 = v72
		goto L40
	} else {
		goto L45
	}
L44:
	;
	v82 = v71
	v83 = v72
	goto L40
L45:
	;
	v75 = int32(1)
	if v71 == v72 {
		v67 = v67 + v75
		v68 = v68 + v75
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(14)
	goto L20
L48:
	;
	goto L49
L49:
	;
	v89 = int32(536124)
	v92 = int32(*(*uint8)(unsafe.Add(mBase, _consts[602])))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v93 == int32(0) {
		v112 = v92
		v113 = v93
		goto L51
	} else {
		goto L52
	}
L50:
	;
	if v113-v112 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L51:
	;
	goto L50
L52:
	;
	if v92 != v93 {
		v112 = v92
		v113 = v93
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v97 = v53
	v98 = v89
	goto L54
L54:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+1)))
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+1)))
	if v102 == int32(0) {
		v112 = v101
		v113 = v102
		goto L51
	} else {
		goto L56
	}
L55:
	;
	v112 = v101
	v113 = v102
	goto L51
L56:
	;
	v105 = int32(1)
	if v101 == v102 {
		v97 = v97 + v105
		v98 = v98 + v105
		goto L54
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(15)
	goto L20
L59:
	;
	goto L60
L60:
	;
	v119 = int32(527504)
	v122 = int32(*(*uint8)(unsafe.Add(mBase, _consts[603])))
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v123 == int32(0) {
		v142 = v122
		v143 = v123
		goto L62
	} else {
		goto L63
	}
L61:
	;
	if v143-v142 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L62:
	;
	goto L61
L63:
	;
	if v122 != v123 {
		v142 = v122
		v143 = v123
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v127 = v53
	v128 = v119
	goto L65
L65:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+1)))
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+1)))
	if v132 == int32(0) {
		v142 = v131
		v143 = v132
		goto L62
	} else {
		goto L67
	}
L66:
	;
	v142 = v131
	v143 = v132
	goto L62
L67:
	;
	v135 = int32(1)
	if v131 == v132 {
		v127 = v127 + v135
		v128 = v128 + v135
		goto L65
	} else {
		goto L68
	}
L68:
	;
	goto L66
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(17)
	goto L20
L70:
	;
	goto L71
L71:
	;
	v149 = int32(542196)
	v152 = int32(*(*uint8)(unsafe.Add(mBase, _consts[604])))
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v153 == int32(0) {
		v172 = v152
		v173 = v153
		goto L73
	} else {
		goto L74
	}
L72:
	;
	if v173-v172 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L73:
	;
	goto L72
L74:
	;
	if v152 != v153 {
		v172 = v152
		v173 = v153
		goto L73
	} else {
		goto L75
	}
L75:
	;
	v157 = v53
	v158 = v149
	goto L76
L76:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+1)))
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+1)))
	if v162 == int32(0) {
		v172 = v161
		v173 = v162
		goto L73
	} else {
		goto L78
	}
L77:
	;
	v172 = v161
	v173 = v162
	goto L73
L78:
	;
	v165 = int32(1)
	if v161 == v162 {
		v157 = v157 + v165
		v158 = v158 + v165
		goto L76
	} else {
		goto L79
	}
L79:
	;
	goto L77
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(18)
	goto L20
L81:
	;
	goto L82
L82:
	;
	v179 = int32(536464)
	v182 = int32(*(*uint8)(unsafe.Add(mBase, _consts[605])))
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v183 == int32(0) {
		v202 = v182
		v203 = v183
		goto L84
	} else {
		goto L85
	}
L83:
	;
	if v203-v202 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L84:
	;
	goto L83
L85:
	;
	if v182 != v183 {
		v202 = v182
		v203 = v183
		goto L84
	} else {
		goto L86
	}
L86:
	;
	v187 = v53
	v188 = v179
	goto L87
L87:
	;
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+1)))
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+1)))
	if v192 == int32(0) {
		v202 = v191
		v203 = v192
		goto L84
	} else {
		goto L89
	}
L88:
	;
	v202 = v191
	v203 = v192
	goto L84
L89:
	;
	v195 = int32(1)
	if v191 == v192 {
		v187 = v187 + v195
		v188 = v188 + v195
		goto L87
	} else {
		goto L90
	}
L90:
	;
	goto L88
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(19)
	goto L20
L92:
	;
	goto L93
L93:
	;
	v209 = int32(524933)
	v212 = int32(*(*uint8)(unsafe.Add(mBase, _consts[606])))
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v213 == int32(0) {
		v232 = v212
		v233 = v213
		goto L95
	} else {
		goto L96
	}
L94:
	;
	if v233-v232 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L95:
	;
	goto L94
L96:
	;
	if v212 != v213 {
		v232 = v212
		v233 = v213
		goto L95
	} else {
		goto L97
	}
L97:
	;
	v217 = v53
	v218 = v209
	goto L98
L98:
	;
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+1)))
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217)+1)))
	if v222 == int32(0) {
		v232 = v221
		v233 = v222
		goto L95
	} else {
		goto L100
	}
L99:
	;
	v232 = v221
	v233 = v222
	goto L95
L100:
	;
	v225 = int32(1)
	if v221 == v222 {
		v217 = v217 + v225
		v218 = v218 + v225
		goto L98
	} else {
		goto L101
	}
L101:
	;
	goto L99
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(21)
	goto L20
L103:
	;
	goto L104
L104:
	;
	v239 = int32(533877)
	v242 = int32(*(*uint8)(unsafe.Add(mBase, _consts[607])))
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v243 == int32(0) {
		v262 = v242
		v263 = v243
		goto L106
	} else {
		goto L107
	}
L105:
	;
	if v263-v262 == int32(0) {
		goto L113
	} else {
		goto L114
	}
L106:
	;
	goto L105
L107:
	;
	if v242 != v243 {
		v262 = v242
		v263 = v243
		goto L106
	} else {
		goto L108
	}
L108:
	;
	v247 = v53
	v248 = v239
	goto L109
L109:
	;
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248)+1)))
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v247)+1)))
	if v252 == int32(0) {
		v262 = v251
		v263 = v252
		goto L106
	} else {
		goto L111
	}
L110:
	;
	v262 = v251
	v263 = v252
	goto L106
L111:
	;
	v255 = int32(1)
	if v251 == v252 {
		v247 = v247 + v255
		v248 = v248 + v255
		goto L109
	} else {
		goto L112
	}
L112:
	;
	goto L110
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(22)
	goto L20
L114:
	;
	goto L115
L115:
	;
	v269 = int32(544320)
	v272 = int32(*(*uint8)(unsafe.Add(mBase, _consts[608])))
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v273 == int32(0) {
		v292 = v272
		v293 = v273
		goto L117
	} else {
		goto L118
	}
L116:
	;
	if v293-v292 == int32(0) {
		goto L124
	} else {
		goto L125
	}
L117:
	;
	goto L116
L118:
	;
	if v272 != v273 {
		v292 = v272
		v293 = v273
		goto L117
	} else {
		goto L119
	}
L119:
	;
	v277 = v53
	v278 = v269
	goto L120
L120:
	;
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v278)+1)))
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v277)+1)))
	if v282 == int32(0) {
		v292 = v281
		v293 = v282
		goto L117
	} else {
		goto L122
	}
L121:
	;
	v292 = v281
	v293 = v282
	goto L117
L122:
	;
	v285 = int32(1)
	if v281 == v282 {
		v277 = v277 + v285
		v278 = v278 + v285
		goto L120
	} else {
		goto L123
	}
L123:
	;
	goto L121
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(23)
	goto L20
L125:
	;
	goto L126
L126:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L12
	} else {
		goto L127
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v53
	F_errmsg_internal(m, int32(724395), v8+int32(16))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L12
	} else {
		goto L128
	}
L128:
	;
	F_errfinish(m, int32(495555), int32(272), int32(417986))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L12
	} else {
		goto L129
	}
L129:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L130:
	;
	v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	v318 = int32(16)
	v320 = int32(63)
	v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+1)))
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+2)))
	v338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+3)))
	v346 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = (v317+v318)&v320 | (v322+v318)&v320<<(uint(int32(6))%32) | (v330+v318)&v320<<(uint(int32(12))%32) | (v338+v318)&v320<<(uint(int32(18))%32) | (v346+v318)&v320<<(uint(int32(24))%32)
	goto L20
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v355
	goto L20
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = v358
	goto L20
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+80)) = v361
	goto L20
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+84)) = v364
	goto L20
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+88)) = v367
	goto L20
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = v370
	goto L20
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+60)) = v373
	goto L20
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+64)) = v376
	goto L20
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+68)) = v379
	goto L20
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+72)) = v382
	goto L20
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+76)) = v385
	goto L20
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v388
	goto L20
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v391
	goto L20
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v394
	goto L20
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v56
	F_errmsg_internal(m, int32(486797), v8)
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L12
	} else {
		goto L146
	}
L146:
	;
	F_errfinish(m, int32(495555), int32(326), int32(417986))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L12
	} else {
		goto L147
	}
L147:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v410
	goto L20
L149:
	;
	v416 = v413 << (uint(int32(24)) % 32)
	if v416 != 0 {
		v52 = v416
		goto L17
	} else {
		goto L150
	}
L150:
	;
	goto L18
L151:
	;
	m.G0 = v8 + int32(48)
	return
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v53
	F_errmsg_internal(m, int32(727661), v8+int32(32))
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L12
	} else {
		goto L153
	}
L153:
	;
	F_errfinish(m, int32(495555), int32(276), int32(417986))
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L12
	} else {
		goto L154
	}
L154:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pq_redirect_to_shm_mq(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	*(*int32)(unsafe.Add(mBase, _consts[599])) = l1
	*(*int32)(unsafe.Add(mBase, _consts[323])) = int32(1618392)
	*(*int32)(unsafe.Add(mBase, _consts[225])) = int32(2)
	*(*int32)(unsafe.Add(mBase, _consts[600])) = int32(196610)
	F_on_dsm_detach(m, l0, int32(808), int32(0))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		return
	}
}
func F_pq_send_ascii_string(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v5 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v7 = l1
	v8 = v5
	goto L4
L2:
	;
	goto L3
L3:
	;
	F_appendStringInfoChar(m, l0, int32(0))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L13
	} else {
		goto L16
	}
L4:
	;
	if base.I32_extend8_s(v8) < int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	v14 = int32(63)
	goto L8
L7:
	;
	v14 = v8
	goto L8
L8:
	;
	v15 = base.I32_extend8_s(v14)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v16 <= v17+int32(1) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v36 = v7 + int32(1)
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
	if v37 != 0 {
		v7 = v36
		v8 = v37
		goto L4
	} else {
		goto L15
	}
L10:
	;
	F_appendStringInfoChar(m, l0, v15)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*uint8)(unsafe.Add(mBase, uint32(v23+v17))) = uint8(v15)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v28 = v26 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v28
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v32 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v30+v28))) = uint8(v32)
	goto L9
L13:
	;
	return
L14:
	;
	goto L9
L15:
	;
	goto L5
L16:
	;
	return
}
