package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_comparetup_datum(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
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
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v9 == int32(1) {
		if v8&int32(1) != 0 {
			v43 = v7
			v46 = int32(0)
			v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)+24))
			if v47 == v46 {
				v84 = v46
				return v84
			} else {
				v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
				v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
				if v51 == int32(1) {
					if v50&int32(1) != 0 {
						v84 = v46
						return v84
					} else {
						v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+9)))
						if v58 != 0 {
							v59 = int32(-1)
						} else {
							v59 = int32(1)
						}
						return v59
					}
				} else {
					if v50&int32(1) != 0 {
						v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+9)))
						if v65 != 0 {
							v66 = int32(1)
						} else {
							v66 = int32(-1)
						}
						return v66
					} else {
						v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						v70 = *(*int32)(unsafe.Add(mBase, uint32(v43)+32))
						v71 = m.T0[v70].(func(*base.Module, int32, int32, int32) int32)(m, v68, v69, v43)
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
							return int32(0)
						} else {
							v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+8)))
							if v73 != int32(1) {
								v84 = v71
							} else {
								v77 = int32(0)
								if v71 < v77 {
									v81 = int32(1)
								} else {
									v81 = v77 - v71
								}
								v84 = v81
							}
							return v84
						}
					}
				}
			}
		} else {
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+9)))
			if v16 != 0 {
				v17 = int32(-1)
			} else {
				v17 = int32(1)
			}
			return v17
		}
	} else {
		if v8&int32(1) != 0 {
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+9)))
			if v23 != 0 {
				v24 = int32(1)
			} else {
				v24 = int32(-1)
			}
			return v24
		} else {
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
			v29 = m.T0[v28].(func(*base.Module, int32, int32, int32) int32)(m, v26, v27, v7)
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int32(0)
			} else {
				v33 = int32(1)
				v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+8)))
				if v34 != v33 {
					v41 = v29
					if v41 != 0 {
						v84 = v41
						return v84
					} else {
						v42 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
						v43 = v42
						v46 = int32(0)
						v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)+24))
						if v47 == v46 {
							v84 = v46
							return v84
						} else {
							v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
							v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
							if v51 == int32(1) {
								if v50&int32(1) != 0 {
									v84 = v46
									return v84
								} else {
									v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+9)))
									if v58 != 0 {
										v59 = int32(-1)
									} else {
										v59 = int32(1)
									}
									return v59
								}
							} else {
								if v50&int32(1) != 0 {
									v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+9)))
									if v65 != 0 {
										v66 = int32(1)
									} else {
										v66 = int32(-1)
									}
									return v66
								} else {
									v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
									v70 = *(*int32)(unsafe.Add(mBase, uint32(v43)+32))
									v71 = m.T0[v70].(func(*base.Module, int32, int32, int32) int32)(m, v68, v69, v43)
									mBase = m.M
									v72 = m.ExcPending
									if v72 != 0 {
										return int32(0)
									} else {
										v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+8)))
										if v73 != int32(1) {
											v84 = v71
										} else {
											v77 = int32(0)
											if v71 < v77 {
												v81 = int32(1)
											} else {
												v81 = v77 - v71
											}
											v84 = v81
										}
										return v84
									}
								}
							}
						}
					}
				} else {
					if v29 < int32(0) {
						v84 = v33
						return v84
					} else {
						v41 = int32(0) - v29
						if v41 != 0 {
							v84 = v41
							return v84
						} else {
							v42 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
							v43 = v42
							v46 = int32(0)
							v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)+24))
							if v47 == v46 {
								v84 = v46
								return v84
							} else {
								v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
								v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
								if v51 == int32(1) {
									if v50&int32(1) != 0 {
										v84 = v46
										return v84
									} else {
										v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+9)))
										if v58 != 0 {
											v59 = int32(-1)
										} else {
											v59 = int32(1)
										}
										return v59
									}
								} else {
									if v50&int32(1) != 0 {
										v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+9)))
										if v65 != 0 {
											v66 = int32(1)
										} else {
											v66 = int32(-1)
										}
										return v66
									} else {
										v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
										v70 = *(*int32)(unsafe.Add(mBase, uint32(v43)+32))
										v71 = m.T0[v70].(func(*base.Module, int32, int32, int32) int32)(m, v68, v69, v43)
										mBase = m.M
										v72 = m.ExcPending
										if v72 != 0 {
											return int32(0)
										} else {
											v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+8)))
											if v73 != int32(1) {
												v84 = v71
											} else {
												v77 = int32(0)
												if v71 < v77 {
													v81 = int32(1)
												} else {
													v81 = v77 - v71
												}
												v84 = v81
											}
											return v84
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_comparetup_datum_tiebreak(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	v4 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+24))
	if v7 == v4 {
		v44 = v4
		return v44
	} else {
		v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
		v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
		if v11 == int32(1) {
			if v10&int32(1) != 0 {
				v44 = v4
				return v44
			} else {
				v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+9)))
				if v18 != 0 {
					v19 = int32(-1)
				} else {
					v19 = int32(1)
				}
				return v19
			}
		} else {
			if v10&int32(1) != 0 {
				v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+9)))
				if v25 != 0 {
					v26 = int32(1)
				} else {
					v26 = int32(-1)
				}
				return v26
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v6)+32))
				v31 = m.T0[v30].(func(*base.Module, int32, int32, int32) int32)(m, v28, v29, v6)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+8)))
					if v35 != int32(1) {
						v44 = v31
					} else {
						v39 = int32(0)
						if v31 < v39 {
							v43 = int32(1)
						} else {
							v43 = v39 - v31
						}
						v44 = v43
					}
					return v44
				}
			}
		}
	}
}
func F_comparetup_index_btree(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v9 == int32(1) {
		if v7&int32(1) != 0 {
			v44 = F_comparetup_index_btree_tiebreak(m, l0, l1, l2)
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return int32(0)
			} else {
				v46 = v44
				return v46
			}
		} else {
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+9)))
			if v16 != 0 {
				v17 = int32(-1)
			} else {
				v17 = int32(1)
			}
			return v17
		}
	} else {
		if v7&int32(1) != 0 {
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+9)))
			if v23 != 0 {
				v24 = int32(1)
			} else {
				v24 = int32(-1)
			}
			return v24
		} else {
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
			v29 = m.T0[v28].(func(*base.Module, int32, int32, int32) int32)(m, v26, v27, v8)
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int32(0)
			} else {
				v33 = int32(1)
				v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+8)))
				if v34 != v33 {
					v41 = v29
					if v41 != 0 {
						v46 = v41
						return v46
					} else {
						v44 = F_comparetup_index_btree_tiebreak(m, l0, l1, l2)
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int32(0)
						} else {
							v46 = v44
							return v46
						}
					}
				} else {
					if v29 < int32(0) {
						v46 = v33
						return v46
					} else {
						v41 = int32(0) - v29
						if v41 != 0 {
							v46 = v41
							return v46
						} else {
							v44 = F_comparetup_index_btree_tiebreak(m, l0, l1, l2)
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return int32(0)
							} else {
								v46 = v44
								return v46
							}
						}
					}
				}
			}
		}
	}
}
func F_comparetup_index_btree_tiebreak(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
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
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
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
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v230 int32
	_ = v230
	v14 = m.G0
	v16 = v14 - int32(208)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l2)+60))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+52))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+24))
	if v25 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v16 + int32(208)
	return v230
L2:
	;
	v71 = int32(2)
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v71 <= v18 {
		goto L27
	} else {
		goto L28
	}
L3:
	;
	v31 = F_index_getattr_2(m, v20, int32(1), v23, v16+int32(207))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	v38 = F_index_getattr_2(m, v19, int32(1), v23, v16+int32(206))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+206)))
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+207)))
	if v41 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	if v40&int32(1) != 0 {
		goto L2
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	if v40&int32(1) != 0 {
		goto L14
	} else {
		goto L15
	}
L10:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+9)))
	if v48 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v49 = int32(-1)
	goto L13
L12:
	;
	v49 = int32(1)
	goto L13
L13:
	;
	v230 = v49
	goto L1
L14:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+9)))
	if v54 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v24)+32))
	v57 = m.T0[v56].(func(*base.Module, int32, int32, int32) int32)(m, v31, v38, v24)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L4
	} else {
		goto L20
	}
L17:
	;
	v55 = int32(1)
	goto L19
L18:
	;
	v55 = int32(-1)
	goto L19
L19:
	;
	v230 = v55
	goto L1
L20:
	;
	v59 = int32(1)
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+8)))
	if v60 != v59 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v67 = v57
	goto L23
L22:
	;
	if v57 < int32(0) {
		v230 = v59
		goto L1
	} else {
		goto L24
	}
L23:
	;
	if v67 != 0 {
		v230 = v67
		goto L1
	} else {
		goto L25
	}
L24:
	;
	v67 = int32(0) - v57
	goto L23
L25:
	;
	goto L2
L26:
	;
	v230 = int32(1)
	goto L1
L27:
	;
	v76 = v24
	v77 = v71
	v86 = v72
	goto L30
L28:
	;
	v142 = v72
	goto L29
L29:
	;
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+8)))
	if v144 != int32(1) {
		goto L55
	} else {
		goto L56
	}
L30:
	;
	v89 = v76 + int32(36)
	v92 = F_index_getattr_2(m, v20, v77, v23, v16+int32(207))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L4
	} else {
		goto L32
	}
L31:
	;
	v142 = v127
	goto L29
L32:
	;
	v96 = F_index_getattr_2(m, v19, v77, v23, v16+int32(206))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+206)))
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+207)))
	if v99 == int32(1) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v129 = v77 + int32(1)
	if v129 <= v18 {
		v76 = v89
		v77 = v129
		v86 = v127
		goto L30
	} else {
		goto L54
	}
L35:
	;
	v102 = int32(1)
	if v98&v102 != 0 {
		v127 = v102
		goto L34
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	if v98&int32(1) != 0 {
		goto L42
	} else {
		goto L43
	}
L38:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+45)))
	if v107 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v108 = int32(-1)
	goto L41
L40:
	;
	v108 = int32(1)
	goto L41
L41:
	;
	v230 = v108
	goto L1
L42:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+45)))
	if v113 != 0 {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	goto L44
L44:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v76)+52))
	v116 = m.T0[v115].(func(*base.Module, int32, int32, int32) int32)(m, v92, v96, v89)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L4
	} else {
		goto L48
	}
L45:
	;
	v114 = int32(1)
	goto L47
L46:
	;
	v114 = int32(-1)
	goto L47
L47:
	;
	v230 = v114
	goto L1
L48:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+44)))
	if v118 == int32(1) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	if v116 < int32(0) {
		goto L26
	} else {
		goto L52
	}
L50:
	;
	v125 = v116
	goto L51
L51:
	;
	if v125 != 0 {
		v230 = v125
		goto L1
	} else {
		goto L53
	}
L52:
	;
	v125 = int32(0) - v116
	goto L51
L53:
	;
	v127 = v86
	goto L34
L54:
	;
	goto L31
L55:
	;
	v206 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+2)))
	v207 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20))))
	v208 = int32(16)
	v210 = v206 | v207<<(uint(v208)%32)
	v211 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19)+2)))
	v212 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19))))
	v215 = v211 | v212<<(uint(v208)%32)
	if v210 != v215 {
		goto L74
	} else {
		goto L75
	}
L56:
	;
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+9)))
	if v147 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v150 = int32(1)
	if (v142^v150)&v150 == int32(0) {
		goto L55
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	F_index_deform_tuple(m, v20, v23, v16-int32(-64), v16+int32(32))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L4
	} else {
		goto L61
	}
L60:
	;
	goto L59
L61:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v167 = F_BuildIndexValueDescription(m, v162, v16-int32(-64), v16+int32(32))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L4
	} else {
		goto L62
	}
L62:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L4
	} else {
		goto L63
	}
L63:
	;
	F_errcode(m, int32(83906754))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L4
	} else {
		goto L64
	}
L64:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v176)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v177 + int32(4)
	F_errmsg(m, int32(696896), v16+int32(16))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L4
	} else {
		goto L65
	}
L65:
	;
	if v167 != 0 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v195)+48))
	F_errtableconstraint(m, v194, v196+int32(4))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L4
	} else {
		goto L72
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v167
	F_errdetail(m, int32(648385), v16)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L4
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	F_errdetail(m, int32(578513), int32(0))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L4
	} else {
		goto L71
	}
L70:
	;
	goto L66
L71:
	;
	goto L66
L72:
	;
	F_errfinish(m, int32(493745), int32(1693), int32(319678))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L4
	} else {
		goto L73
	}
L73:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L74:
	;
	if base.Ui32(v210) < base.Ui32(v215) {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	goto L76
L76:
	;
	v221 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+4)))
	v222 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19)+4)))
	v230 = base.B2i32(base.Ui32(v222) < base.Ui32(v221)) - base.B2i32(base.Ui32(v221) < base.Ui32(v222))
	goto L1
L77:
	;
	v220 = int32(-1)
	goto L79
L78:
	;
	v220 = int32(1)
	goto L79
L79:
	;
	v230 = v220
	goto L1
}
