package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_tsquery_numnode(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+4))
	return v3
}
func F_tsquery_rewrite(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = F_pg_detoast_datum_copy(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
		if v18 == int32(0) {
			v99 = v14
			m.G0 = v11 + int32(16)
			return v99
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
			if v22 == int32(0) {
				v99 = v14
				m.G0 = v11 + int32(16)
				return v99
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				v27 = v14 + int32(8)
				v31 = F_QT2QTN(m, v27, v27+v18*int32(12))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					F_QTNTernary(m, v31)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						F_QTNSort(m, v31)
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int32(0)
						} else {
							v38 = v21 + int32(8)
							v39 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
							v43 = F_QT2QTN(m, v38, v38+v39*int32(12))
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return int32(0)
							} else {
								F_QTNTernary(m, v43)
								mBase = m.M
								v46 = m.ExcPending
								if v46 != 0 {
									return int32(0)
								} else {
									F_QTNSort(m, v43)
									mBase = m.M
									v48 = m.ExcPending
									if v48 != 0 {
										return int32(0)
									} else {
										v49 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
										if v49 != 0 {
											v51 = v25 + int32(8)
											v55 = F_QT2QTN(m, v51, v51+v49*int32(12))
											mBase = m.M
											v56 = m.ExcPending
											if v56 != 0 {
												return int32(0)
											} else {
												v59 = v55
												v62 = F_dofindsubquery(m, v31, v43, v59, v11+int32(15))
												mBase = m.M
												v63 = m.ExcPending
												if v63 != 0 {
													return int32(0)
												} else {
													F_QTNFree(m, v43)
													mBase = m.M
													v65 = m.ExcPending
													if v65 != 0 {
														return int32(0)
													} else {
														F_QTNFree(m, v59)
														mBase = m.M
														v67 = m.ExcPending
														if v67 != 0 {
															return int32(0)
														} else {
															if v62 == int32(0) {
																*(*int64)(unsafe.Add(mBase, uint32(v14))) = int64(32)
																v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
																if v72 != v21 {
																	F_pfree(m, v21)
																	mBase = m.M
																	v75 = m.ExcPending
																	if v75 != 0 {
																		return int32(0)
																	} else {
																		v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
																		if v25 != v76 {
																			v94 = v14
																			F_pfree(m, v25)
																			mBase = m.M
																			v96 = m.ExcPending
																			if v96 != 0 {
																				return int32(0)
																			} else {
																				v99 = v94
																				m.G0 = v11 + int32(16)
																				return v99
																			}
																		} else {
																			v99 = v14
																			m.G0 = v11 + int32(16)
																			return v99
																		}
																	}
																} else {
																	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
																	if v25 != v76 {
																		v94 = v14
																		F_pfree(m, v25)
																		mBase = m.M
																		v96 = m.ExcPending
																		if v96 != 0 {
																			return int32(0)
																		} else {
																			v99 = v94
																			m.G0 = v11 + int32(16)
																			return v99
																		}
																	} else {
																		v99 = v14
																		m.G0 = v11 + int32(16)
																		return v99
																	}
																}
															} else {
																F_QTNBinary(m, v62)
																mBase = m.M
																v79 = m.ExcPending
																if v79 != 0 {
																	return int32(0)
																} else {
																	v80 = F_QTN2QT(m, v62)
																	mBase = m.M
																	v81 = m.ExcPending
																	if v81 != 0 {
																		return int32(0)
																	} else {
																		F_QTNFree(m, v62)
																		mBase = m.M
																		v83 = m.ExcPending
																		if v83 != 0 {
																			return int32(0)
																		} else {
																			v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																			if v84 != v14 {
																				F_pfree(m, v14)
																				mBase = m.M
																				v87 = m.ExcPending
																				if v87 != 0 {
																					return int32(0)
																				} else {
																					v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
																					if v88 != v21 {
																						F_pfree(m, v21)
																						mBase = m.M
																						v91 = m.ExcPending
																						if v91 != 0 {
																							return int32(0)
																						} else {
																							v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
																							if v25 == v92 {
																								v99 = v80
																								m.G0 = v11 + int32(16)
																								return v99
																							} else {
																								v94 = v80
																								F_pfree(m, v25)
																								mBase = m.M
																								v96 = m.ExcPending
																								if v96 != 0 {
																									return int32(0)
																								} else {
																									v99 = v94
																									m.G0 = v11 + int32(16)
																									return v99
																								}
																							}
																						}
																					} else {
																						v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
																						if v25 == v92 {
																							v99 = v80
																							m.G0 = v11 + int32(16)
																							return v99
																						} else {
																							v94 = v80
																							F_pfree(m, v25)
																							mBase = m.M
																							v96 = m.ExcPending
																							if v96 != 0 {
																								return int32(0)
																							} else {
																								v99 = v94
																								m.G0 = v11 + int32(16)
																								return v99
																							}
																						}
																					}
																				}
																			} else {
																				v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
																				if v88 != v21 {
																					F_pfree(m, v21)
																					mBase = m.M
																					v91 = m.ExcPending
																					if v91 != 0 {
																						return int32(0)
																					} else {
																						v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
																						if v25 == v92 {
																							v99 = v80
																							m.G0 = v11 + int32(16)
																							return v99
																						} else {
																							v94 = v80
																							F_pfree(m, v25)
																							mBase = m.M
																							v96 = m.ExcPending
																							if v96 != 0 {
																								return int32(0)
																							} else {
																								v99 = v94
																								m.G0 = v11 + int32(16)
																								return v99
																							}
																						}
																					}
																				} else {
																					v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
																					if v25 == v92 {
																						v99 = v80
																						m.G0 = v11 + int32(16)
																						return v99
																					} else {
																						v94 = v80
																						F_pfree(m, v25)
																						mBase = m.M
																						v96 = m.ExcPending
																						if v96 != 0 {
																							return int32(0)
																						} else {
																							v99 = v94
																							m.G0 = v11 + int32(16)
																							return v99
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
											}
										} else {
											v59 = int32(0)
											v62 = F_dofindsubquery(m, v31, v43, v59, v11+int32(15))
											mBase = m.M
											v63 = m.ExcPending
											if v63 != 0 {
												return int32(0)
											} else {
												F_QTNFree(m, v43)
												mBase = m.M
												v65 = m.ExcPending
												if v65 != 0 {
													return int32(0)
												} else {
													F_QTNFree(m, v59)
													mBase = m.M
													v67 = m.ExcPending
													if v67 != 0 {
														return int32(0)
													} else {
														if v62 == int32(0) {
															*(*int64)(unsafe.Add(mBase, uint32(v14))) = int64(32)
															v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
															if v72 != v21 {
																F_pfree(m, v21)
																mBase = m.M
																v75 = m.ExcPending
																if v75 != 0 {
																	return int32(0)
																} else {
																	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
																	if v25 != v76 {
																		v94 = v14
																		F_pfree(m, v25)
																		mBase = m.M
																		v96 = m.ExcPending
																		if v96 != 0 {
																			return int32(0)
																		} else {
																			v99 = v94
																			m.G0 = v11 + int32(16)
																			return v99
																		}
																	} else {
																		v99 = v14
																		m.G0 = v11 + int32(16)
																		return v99
																	}
																}
															} else {
																v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
																if v25 != v76 {
																	v94 = v14
																	F_pfree(m, v25)
																	mBase = m.M
																	v96 = m.ExcPending
																	if v96 != 0 {
																		return int32(0)
																	} else {
																		v99 = v94
																		m.G0 = v11 + int32(16)
																		return v99
																	}
																} else {
																	v99 = v14
																	m.G0 = v11 + int32(16)
																	return v99
																}
															}
														} else {
															F_QTNBinary(m, v62)
															mBase = m.M
															v79 = m.ExcPending
															if v79 != 0 {
																return int32(0)
															} else {
																v80 = F_QTN2QT(m, v62)
																mBase = m.M
																v81 = m.ExcPending
																if v81 != 0 {
																	return int32(0)
																} else {
																	F_QTNFree(m, v62)
																	mBase = m.M
																	v83 = m.ExcPending
																	if v83 != 0 {
																		return int32(0)
																	} else {
																		v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																		if v84 != v14 {
																			F_pfree(m, v14)
																			mBase = m.M
																			v87 = m.ExcPending
																			if v87 != 0 {
																				return int32(0)
																			} else {
																				v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
																				if v88 != v21 {
																					F_pfree(m, v21)
																					mBase = m.M
																					v91 = m.ExcPending
																					if v91 != 0 {
																						return int32(0)
																					} else {
																						v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
																						if v25 == v92 {
																							v99 = v80
																							m.G0 = v11 + int32(16)
																							return v99
																						} else {
																							v94 = v80
																							F_pfree(m, v25)
																							mBase = m.M
																							v96 = m.ExcPending
																							if v96 != 0 {
																								return int32(0)
																							} else {
																								v99 = v94
																								m.G0 = v11 + int32(16)
																								return v99
																							}
																						}
																					}
																				} else {
																					v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
																					if v25 == v92 {
																						v99 = v80
																						m.G0 = v11 + int32(16)
																						return v99
																					} else {
																						v94 = v80
																						F_pfree(m, v25)
																						mBase = m.M
																						v96 = m.ExcPending
																						if v96 != 0 {
																							return int32(0)
																						} else {
																							v99 = v94
																							m.G0 = v11 + int32(16)
																							return v99
																						}
																					}
																				}
																			}
																		} else {
																			v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
																			if v88 != v21 {
																				F_pfree(m, v21)
																				mBase = m.M
																				v91 = m.ExcPending
																				if v91 != 0 {
																					return int32(0)
																				} else {
																					v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
																					if v25 == v92 {
																						v99 = v80
																						m.G0 = v11 + int32(16)
																						return v99
																					} else {
																						v94 = v80
																						F_pfree(m, v25)
																						mBase = m.M
																						v96 = m.ExcPending
																						if v96 != 0 {
																							return int32(0)
																						} else {
																							v99 = v94
																							m.G0 = v11 + int32(16)
																							return v99
																						}
																					}
																				}
																			} else {
																				v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
																				if v25 == v92 {
																					v99 = v80
																					m.G0 = v11 + int32(16)
																					return v99
																				} else {
																					v94 = v80
																					F_pfree(m, v25)
																					mBase = m.M
																					v96 = m.ExcPending
																					if v96 != 0 {
																						return int32(0)
																					} else {
																						v99 = v94
																						m.G0 = v11 + int32(16)
																						return v99
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
