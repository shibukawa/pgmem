package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AsyncExistsPendingNotify(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	v2 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = l0
	v18 = *(*int32)(unsafe.Add(mBase, _consts[247]))
	if v18 == v2 {
		v145 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v14 + int32(16)
	return v145
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	if v21 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v24 = int32(0)
	v26 = F_hash_search(m, v21, v14+int32(12), v24, v24)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v33 == int32(0) {
		v145 = v2
		goto L1
	} else {
		goto L9
	}
L6:
	;
	return int32(0)
L7:
	;
	if v26 == int32(0) {
		v145 = v2
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v145 = int32(1)
	goto L1
L9:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	if v36 <= int32(0) {
		v145 = v2
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v39 = int32(0)
	if v39 < v36 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v43 = v36
	goto L13
L12:
	;
	v43 = v39
	goto L13
L13:
	;
	v45 = l0 + int32(4)
	v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	v52 = v39
	goto L14
L14:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v49+v52<<(uint(int32(2))%32))))
	v65 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v64))))
	if v46 != v65 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v145 = v2
	goto L1
L16:
	;
	v138 = v52 + int32(1)
	if v138 != v43 {
		v52 = v138
		goto L14
	} else {
		goto L38
	}
L17:
	;
	v67 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+2)))
	v68 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v64)+2)))
	if v67 != v68 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v70 = int32(4)
	v71 = v64 + v70
	v72 = v46 + int32(2) + v67
	if base.Ui32(v70) <= base.Ui32(v72) {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	if v134 != 0 {
		goto L16
	} else {
		goto L37
	}
L20:
	;
	v134 = int32(0)
	goto L19
L21:
	;
	v108 = v103
	v109 = v104
	v110 = v105
	goto L31
L22:
	;
	if (v45|v71)&int32(3) != 0 {
		v103 = v45
		v104 = v71
		v105 = v72
		goto L21
	} else {
		goto L25
	}
L23:
	;
	v96 = v45
	v97 = v71
	v98 = v72
	goto L24
L24:
	;
	if v98 == int32(0) {
		goto L20
	} else {
		goto L30
	}
L25:
	;
	v80 = v45
	v81 = v71
	v82 = v72
	goto L26
L26:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	if v85 != v86 {
		v103 = v80
		v104 = v81
		v105 = v82
		goto L21
	} else {
		goto L28
	}
L27:
	;
	v96 = v91
	v97 = v89
	v98 = v93
	goto L24
L28:
	;
	v88 = int32(4)
	v89 = v81 + v88
	v91 = v80 + v88
	v93 = v82 - v88
	if base.Ui32(int32(3)) < base.Ui32(v93) {
		v80 = v91
		v81 = v89
		v82 = v93
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v103 = v96
	v104 = v97
	v105 = v98
	goto L21
L31:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108))))
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109))))
	if v113 == v114 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v134 = v113 - v114
	goto L19
L33:
	;
	v116 = int32(1)
	v121 = v110 - v116
	if v121 != 0 {
		v108 = v108 + v116
		v109 = v109 + v116
		v110 = v121
		goto L31
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	goto L32
L36:
	;
	goto L20
L37:
	;
	v145 = int32(1)
	goto L1
L38:
	;
	goto L15
}
func F_ExecAsyncRequest(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 float64
	_ = v42
	var v44 float64
	_ = v44
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+52))
	if v10 != 0 {
		F_ExecReScan(m, v9)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v14 = v13
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
			if v15 != 0 {
				F_InstrStartNode(m, v15)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return
				} else {
					v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v19 = v18
					v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
					if v20 == int32(418) {
						v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+128))
						v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+172))
						m.T0[v25].(func(*base.Module, int32))(m, l0)
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return
						} else {
							v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
							if v29 == int32(397) {
								F_ExecAsyncAppendResponse(m, l0)
								mBase = m.M
								v33 = m.ExcPending
								if v33 != 0 {
									return
								} else {
									v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+20))
									if v35 != 0 {
										v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
										if v36 != 0 {
											v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+4)))
											if v39&int32(2) != 0 {
												v42 = float64(0)
											} else {
												v42 = float64(1)
											}
											v44 = v42
										} else {
											v44 = float64(0)
										}
										F_InstrStopNode(m, v35, v44)
										mBase = m.M
										v46 = m.ExcPending
										if v46 != 0 {
											return
										} else {
											m.G0 = v7 + int32(32)
											return
										}
									} else {
										m.G0 = v7 + int32(32)
										return
									}
								}
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return
								} else {
									v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
									*(*int32)(unsafe.Add(mBase, uint32(v7))) = v56
									F_errmsg_internal(m, int32(482882), v7)
									mBase = m.M
									v60 = m.ExcPending
									if v60 != 0 {
										return
									} else {
										F_errfinish(m, int32(496759), int32(127), int32(359074))
										mBase = m.M
										v65 = m.ExcPending
										if v65 != 0 {
											return
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
						v69 = m.ExcPending
						if v69 != 0 {
							return
						} else {
							v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
							*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v71
							F_errmsg_internal(m, int32(482882), v7+int32(16))
							mBase = m.M
							v77 = m.ExcPending
							if v77 != 0 {
								return
							} else {
								F_errfinish(m, int32(496759), int32(43), int32(76891))
								mBase = m.M
								v82 = m.ExcPending
								if v82 != 0 {
									return
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
				v19 = v14
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
				if v20 == int32(418) {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+128))
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+172))
					m.T0[v25].(func(*base.Module, int32))(m, l0)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
						if v29 == int32(397) {
							F_ExecAsyncAppendResponse(m, l0)
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return
							} else {
								v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+20))
								if v35 != 0 {
									v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									if v36 != 0 {
										v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+4)))
										if v39&int32(2) != 0 {
											v42 = float64(0)
										} else {
											v42 = float64(1)
										}
										v44 = v42
									} else {
										v44 = float64(0)
									}
									F_InstrStopNode(m, v35, v44)
									mBase = m.M
									v46 = m.ExcPending
									if v46 != 0 {
										return
									} else {
										m.G0 = v7 + int32(32)
										return
									}
								} else {
									m.G0 = v7 + int32(32)
									return
								}
							}
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return
							} else {
								v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
								*(*int32)(unsafe.Add(mBase, uint32(v7))) = v56
								F_errmsg_internal(m, int32(482882), v7)
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return
								} else {
									F_errfinish(m, int32(496759), int32(127), int32(359074))
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
										return
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
					v69 = m.ExcPending
					if v69 != 0 {
						return
					} else {
						v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v71
						F_errmsg_internal(m, int32(482882), v7+int32(16))
						mBase = m.M
						v77 = m.ExcPending
						if v77 != 0 {
							return
						} else {
							F_errfinish(m, int32(496759), int32(43), int32(76891))
							mBase = m.M
							v82 = m.ExcPending
							if v82 != 0 {
								return
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
	} else {
		v14 = v9
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
		if v15 != 0 {
			F_InstrStartNode(m, v15)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v19 = v18
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
				if v20 == int32(418) {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+128))
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+172))
					m.T0[v25].(func(*base.Module, int32))(m, l0)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
						if v29 == int32(397) {
							F_ExecAsyncAppendResponse(m, l0)
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return
							} else {
								v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+20))
								if v35 != 0 {
									v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									if v36 != 0 {
										v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+4)))
										if v39&int32(2) != 0 {
											v42 = float64(0)
										} else {
											v42 = float64(1)
										}
										v44 = v42
									} else {
										v44 = float64(0)
									}
									F_InstrStopNode(m, v35, v44)
									mBase = m.M
									v46 = m.ExcPending
									if v46 != 0 {
										return
									} else {
										m.G0 = v7 + int32(32)
										return
									}
								} else {
									m.G0 = v7 + int32(32)
									return
								}
							}
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return
							} else {
								v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
								*(*int32)(unsafe.Add(mBase, uint32(v7))) = v56
								F_errmsg_internal(m, int32(482882), v7)
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return
								} else {
									F_errfinish(m, int32(496759), int32(127), int32(359074))
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
										return
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
					v69 = m.ExcPending
					if v69 != 0 {
						return
					} else {
						v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v71
						F_errmsg_internal(m, int32(482882), v7+int32(16))
						mBase = m.M
						v77 = m.ExcPending
						if v77 != 0 {
							return
						} else {
							F_errfinish(m, int32(496759), int32(43), int32(76891))
							mBase = m.M
							v82 = m.ExcPending
							if v82 != 0 {
								return
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
			v19 = v14
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
			if v20 == int32(418) {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+128))
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+172))
				m.T0[v25].(func(*base.Module, int32))(m, l0)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
					if v29 == int32(397) {
						F_ExecAsyncAppendResponse(m, l0)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return
						} else {
							v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+20))
							if v35 != 0 {
								v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								if v36 != 0 {
									v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+4)))
									if v39&int32(2) != 0 {
										v42 = float64(0)
									} else {
										v42 = float64(1)
									}
									v44 = v42
								} else {
									v44 = float64(0)
								}
								F_InstrStopNode(m, v35, v44)
								mBase = m.M
								v46 = m.ExcPending
								if v46 != 0 {
									return
								} else {
									m.G0 = v7 + int32(32)
									return
								}
							} else {
								m.G0 = v7 + int32(32)
								return
							}
						}
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return
						} else {
							v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = v56
							F_errmsg_internal(m, int32(482882), v7)
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return
							} else {
								F_errfinish(m, int32(496759), int32(127), int32(359074))
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return
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
				v69 = m.ExcPending
				if v69 != 0 {
					return
				} else {
					v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
					*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v71
					F_errmsg_internal(m, int32(482882), v7+int32(16))
					mBase = m.M
					v77 = m.ExcPending
					if v77 != 0 {
						return
					} else {
						F_errfinish(m, int32(496759), int32(43), int32(76891))
						mBase = m.M
						v82 = m.ExcPending
						if v82 != 0 {
							return
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
}
func F_asyncQueueReadAllNotifications(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int64
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int64
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
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
	var v60 int32
	_ = v60
	var v75 int64
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v97 int32
	_ = v97
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int64
	_ = v110
	var v113 int64
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
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
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int64
	_ = v165
	var v166 int64
	_ = v166
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v298 int64
	_ = v298
	var v300 int32
	_ = v300
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v321 int64
	_ = v321
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	v17 = m.G0
	v19 = v17 - int32(8208)
	m.G0 = v19
	v22 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	v26 = F_LWLockAcquire(m, v22+int32(3456), int32(1))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _consts[248]))
	v31 = *(*int32)(unsafe.Add(mBase, _consts[249]))
	v34 = v29 + v31<<(uint(int32(5))%32)
	v35 = *(*int64)(unsafe.Add(mBase, uint32(v34)+72))
	*(*int64)(unsafe.Add(mBase, uint32(v19)+8)) = v35
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v34)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v37
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v34)+84))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	v41 = *(*int64)(unsafe.Add(mBase, uint32(v29)))
	v43 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	F_LWLockRelease(m, v43+int32(3456))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if base.B2i32(v37 == v40)&base.B2i32(v35 == v41) == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v53 = F_GetLatestSnapshot(m)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	m.G0 = v19 + int32(8208)
	return
L7:
	;
	v55 = F_RegisterSnapshot(m, v53)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v57 = int32(4481623)
	v58 = int32(*(*uint8)(unsafe.Add(mBase, _consts[250])))
	v60 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[250])) = uint8(v60)
	v75 = v35
	goto L9
L9:
	;
	v80 = F_SimpleLruReadPage_ReadOnly(m, int32(4383800), v75, int32(0))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	v308 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	v312 = F_LWLockAcquire(m, v308+int32(3456), int32(1))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L1
	} else {
		goto L66
	}
L11:
	;
	v83 = *(*int32)(unsafe.Add(mBase, _consts[251]))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v84+v80<<(uint(int32(2))%32))))
	v97 = v19 + int32(16)
	goto L12
L12:
	;
	v107 = int32(0)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v110 = *(*int64)(unsafe.Add(mBase, uint32(v19)+8))
	if base.B2i32(v40 == v108)&base.B2i32(v41 == v110) != 0 {
		v158 = v97
		v159 = v107
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v162 = *(*int32)(unsafe.Add(mBase, _consts[251]))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v162)+28))
	v165 = int64(*(*uint16)(unsafe.Add(mBase, _consts[252])))
	v166 = base.I64_rem_s(v75, v165)
	F_LWLockRelease(m, v163+base.I32_wrap_i64(v166)<<(uint(int32(7))%32))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L33
	}
L14:
	;
	goto L13
L15:
	;
	v113 = *(*int64)(unsafe.Add(mBase, uint32(v19)+8))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v115 = v108 + v88
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	v117 = v114 + v116
	v119 = v117 - int32(8173)
	v121 = base.B2i32(base.Ui32(v119) < base.Ui32(int32(-8193)))
	*(*int64)(unsafe.Add(mBase, uint32(v19)+8)) = v113 + base.I64_extend_i32_u(v121)
	if base.Ui32(v119) < base.Ui32(int32(-8193)) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v126 = int32(0)
	goto L18
L17:
	;
	v126 = v117
	goto L18
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v126
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	v130 = *(*int32)(unsafe.Add(mBase, _consts[226]))
	if v128 != v130 {
		v152 = v97
		goto L19
	} else {
		goto L20
	}
L19:
	;
	if base.Ui32(int32(-8194)) < base.Ui32(v119) {
		v97 = v152
		goto L12
	} else {
		goto L32
	}
L20:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v115)+8))
	v133 = F_XidInMVCCSnapshot(m, v132, v55)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	if v133 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v108
	*(*int64)(unsafe.Add(mBase, uint32(v19)+8)) = v110
	v158 = v97
	v159 = int32(1)
	goto L14
L23:
	;
	goto L24
L24:
	;
	v139 = *(*int32)(unsafe.Add(mBase, _consts[253]))
	if v139 == int32(0) {
		v152 = v97
		goto L19
	} else {
		goto L25
	}
L25:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v115)+8))
	v143 = F_TransactionIdDidCommit(m, v142)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	if v143 == int32(0) {
		v152 = v97
		goto L19
	} else {
		goto L27
	}
L27:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	if v147 != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	v152 = v149 + v150
	goto L19
L29:
	;
	v148 = F__emscripten_memcpy_bulkmem(m, v97, v115, v147)
	mBase = m.M
	v149 = v148
	goto L31
L30:
	;
	v149 = v97
	goto L31
L31:
	;
	goto L28
L32:
	;
	v158 = v152
	v159 = v107
	goto L14
L33:
	;
	if base.Ui32(v19+int32(16)) < base.Ui32(v158) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v177 = *(*int32)(unsafe.Add(mBase, _consts[253]))
	v183 = v177
	v184 = v19 + int32(16)
	goto L37
L35:
	;
	goto L36
L36:
	;
	v298 = *(*int64)(unsafe.Add(mBase, uint32(v19)+8))
	if v41 == v298 {
		goto L61
	} else {
		goto L62
	}
L37:
	;
	if v183 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	goto L36
L39:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v184)))
	v280 = v184 + v279
	if base.Ui32(v280) < base.Ui32(v158) {
		v183 = v266
		v184 = v280
		goto L37
	} else {
		goto L59
	}
L40:
	;
	v266 = int32(0)
	goto L39
L41:
	;
	goto L42
L42:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v183)+4))
	if v199 <= int32(0) {
		v266 = v183
		goto L39
	} else {
		goto L43
	}
L43:
	;
	v203 = v184 + int32(16)
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v183)+12))
	v207 = int32(0)
	goto L44
L44:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v204+v207<<(uint(int32(2))%32))))
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203))))
	v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225))))
	if v229 == int32(0) {
		v248 = v228
		v249 = v229
		goto L47
	} else {
		goto L48
	}
L45:
	;
	v254 = F_strlen(m, v203)
	mBase = m.M
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v184)+12))
	F_NotifyMyFrontEnd(m, v203, v254+v203+int32(1), v258)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L1
	} else {
		goto L58
	}
L46:
	;
	if v249-v248 != 0 {
		goto L54
	} else {
		goto L55
	}
L47:
	;
	goto L46
L48:
	;
	if v228 != v229 {
		v248 = v228
		v249 = v229
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v233 = v225
	v234 = v203
	goto L50
L50:
	;
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v234)+1)))
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233)+1)))
	if v238 == int32(0) {
		v248 = v237
		v249 = v238
		goto L47
	} else {
		goto L52
	}
L51:
	;
	v248 = v237
	v249 = v238
	goto L47
L52:
	;
	v241 = int32(1)
	if v237 == v238 {
		v233 = v233 + v241
		v234 = v234 + v241
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	v252 = v207 + int32(1)
	if v252 != v199 {
		v207 = v252
		goto L44
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	goto L45
L57:
	;
	v266 = v183
	goto L39
L58:
	;
	v262 = *(*int32)(unsafe.Add(mBase, _consts[253]))
	v266 = v262
	goto L39
L59:
	;
	goto L38
L60:
	;
	goto L10
L61:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v159|base.B2i32(v300 == v40) == int32(0) {
		v75 = v298
		goto L9
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	if v159 == int32(0) {
		v75 = v298
		goto L9
	} else {
		goto L65
	}
L64:
	;
	goto L60
L65:
	;
	goto L60
L66:
	;
	v315 = *(*int32)(unsafe.Add(mBase, _consts[248]))
	v317 = *(*int32)(unsafe.Add(mBase, _consts[249]))
	v320 = v315 + v317<<(uint(int32(5))%32)
	v321 = *(*int64)(unsafe.Add(mBase, uint32(v19)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v320)+72)) = v321
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	*(*int32)(unsafe.Add(mBase, uint32(v320)+84)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v320)+80)) = v323
	v327 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	F_LWLockRelease(m, v327+int32(3456))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	*(*uint8)(unsafe.Add(mBase, _consts[250])) = uint8(v58)
	F_UnregisterSnapshot(m, v55)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	goto L6
}
